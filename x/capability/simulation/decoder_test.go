package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Finschia/finschia-sdk/simapp"
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/types/kv"
	"github.com/Finschia/finschia-sdk/x/capability/simulation"
	"github.com/Finschia/finschia-sdk/x/capability/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	capOwners := types.CapabilityOwners{
		Owners: []types.Owner{{Module: "transfer", Name: "ports/transfer"}},
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   types.KeyIndex,
				Value: sdk.Uint64ToBigEndian(10),
			},
			{
				Key:   types.KeyPrefixIndexCapability,
				Value: cdc.MustMarshal(&capOwners),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Index", "Index A: 10\nIndex B: 10\n"},
		{"CapabilityOwners", fmt.Sprintf("CapabilityOwners A: %v\nCapabilityOwners B: %v\n", capOwners, capOwners)},
		{"other", ""},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}

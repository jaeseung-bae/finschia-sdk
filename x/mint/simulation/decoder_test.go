package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Finschia/finschia-sdk/simapp"
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/types/kv"
	"github.com/Finschia/finschia-sdk/x/mint/simulation"
	"github.com/Finschia/finschia-sdk/x/mint/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	minter := types.NewMinter(sdk.OneDec(), sdk.NewDec(15))

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.MinterKey, Value: cdc.MustMarshal(&minter)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
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

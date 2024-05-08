package types_test

import (
	"testing"

	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/Finschia/finschia-sdk/x/fswap/types"
)

func TestValidateGenesis(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState *types.GenesisState
		wantErr  bool
	}{
		{
			name:     "valid default genesis",
			genState: types.DefaultGenesis(),
			wantErr:  false,
		},
		{
			name: "valid: initial empty genesis",
			genState: &types.GenesisState{
				Swaps: []types.Swap{},
				SwapStats: types.SwapStats{
					SwapCount: 0,
				},
				Swappeds: []types.Swapped{},
			},
			wantErr: false,
		},
		{
			name: "invalid: the number of Swaps should match to the number of Swappeds",
			genState: &types.GenesisState{
				Swaps: []types.Swap{
					{
						FromDenom:           "from",
						ToDenom:             "toD",
						AmountCapForToDenom: sdk.OneInt(),
						SwapRate:            sdk.NewDec(12),
					},
				},
				SwapStats: types.SwapStats{},
				Swappeds:  nil,
			},
			wantErr: true,
		},
		{
			name: "invalid: SwapCount should match to the number of Swaps",
			genState: &types.GenesisState{
				Swaps: []types.Swap{},
				SwapStats: types.SwapStats{
					SwapCount: 1,
				},
				Swappeds: []types.Swapped{},
			},
			wantErr: true,
		},
		{
			name: "invalid: invalid Swap element with zero AmountCapForToDenom",
			genState: &types.GenesisState{
				Swaps: []types.Swap{
					{
						FromDenom:           "from",
						ToDenom:             "toD",
						AmountCapForToDenom: sdk.ZeroInt(),
						SwapRate:            sdk.NewDec(12),
					},
				},
				SwapStats: types.SwapStats{
					SwapCount: 1,
				},
				Swappeds: []types.Swapped{
					{
						FromCoinAmount: sdk.NewCoin("from", sdk.ZeroInt()),
						ToCoinAmount:   sdk.NewCoin("toD", sdk.ZeroInt()),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid: invalid Swap element with zero SwapRate",
			genState: &types.GenesisState{
				Swaps: []types.Swap{
					{
						FromDenom:           "from",
						ToDenom:             "toD",
						AmountCapForToDenom: sdk.OneInt(),
						SwapRate:            sdk.NewDec(0),
					},
				},
				SwapStats: types.SwapStats{
					SwapCount: 1,
				},
				Swappeds: []types.Swapped{
					{
						FromCoinAmount: sdk.NewCoin("from", sdk.ZeroInt()),
						ToCoinAmount:   sdk.NewCoin("toD", sdk.ZeroInt()),
					},
				},
			},
			wantErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

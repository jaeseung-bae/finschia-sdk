package types

import (
	"fmt"
	"testing"

	sdk "github.com/Finschia/finschia-sdk/types"
	banktypes "github.com/Finschia/finschia-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

func TestMakeSwapProposalValidateBasic(t *testing.T) {
	swapAllExpectedBalance, ok := sdk.NewIntFromString("18281438845984584000000")
	require.True(t, ok)
	pebSwapRateForCony, err := sdk.NewDecFromStr("148079656000000")
	require.NoError(t, err)
	metadata := banktypes.Metadata{
		Description: "",
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: "peb", Exponent: 0, Aliases: nil},
		},
		Base:    "peb",
		Display: "peb",
		Name:    "DUMMY",
		Symbol:  "PEB",
	}

	tests := []struct {
		name             string
		makeSwapProposal *MakeSwapProposal
		wantErr          bool
	}{
		{
			name: "valid",
			makeSwapProposal: NewMakeSwapProposal(
				"valid MakeSwapProposal",
				"Valid description",
				Swap{
					FromDenom:           "cony",
					ToDenom:             "peb",
					AmountCapForToDenom: swapAllExpectedBalance,
					SwapRate:            pebSwapRateForCony,
				},
				metadata,
			),
			wantErr: false,
		},
		{
			name: "invalid: Swap with zero AmountCapForToDenom",
			makeSwapProposal: NewMakeSwapProposal(
				"valid MakeSwapProposal",
				"Valid description",
				Swap{
					FromDenom:           "cony",
					ToDenom:             "peb",
					AmountCapForToDenom: sdk.ZeroInt(),
					SwapRate:            pebSwapRateForCony,
				},
				metadata,
			),
			wantErr: true,
		},
		{
			name: "invalid: Swap with zero SwapRate",
			makeSwapProposal: NewMakeSwapProposal(
				"valid MakeSwapProposal",
				"Valid description",
				Swap{
					FromDenom:           "cony",
					ToDenom:             "peb",
					AmountCapForToDenom: swapAllExpectedBalance,
					SwapRate:            sdk.NewDec(0),
				},
				metadata,
			),
			wantErr: true,
		},
		{
			name: "invalid: Swap with mismatched to-denom",
			makeSwapProposal: NewMakeSwapProposal(
				"valid MakeSwapProposal",
				"Valid description",
				Swap{
					FromDenom:           "cony",
					ToDenom:             "mismatched" + "peb",
					AmountCapForToDenom: swapAllExpectedBalance,
					SwapRate:            pebSwapRateForCony,
				},
				metadata,
			),
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.makeSwapProposal.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				fmt.Println(err)
				return
			}
			require.NoError(t, err)
		})
	}
}

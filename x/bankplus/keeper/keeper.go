package keeper

import (
	"context"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

var _ Keeper = (*BaseKeeper)(nil)

type Keeper interface {
	bankkeeper.Keeper

	AddToInactiveAddr(ctx context.Context, address sdk.AccAddress)
	DeleteFromInactiveAddr(ctx context.Context, address sdk.AccAddress)
	IsInactiveAddr(address sdk.AccAddress) bool

	InitializeBankPlus(ctx context.Context)
}

type BaseKeeper struct {
	bankkeeper.BaseKeeper

	ak      types.AccountKeeper
	cdc     codec.Codec
	addrCdc address.Codec

	storeService   store.KVStoreService
	inactiveAddrs  map[string]bool
	deactMultiSend bool
}

func NewBaseKeeper(
	cdc codec.Codec, storeService store.KVStoreService, ak types.AccountKeeper,
	blockedAddr map[string]bool, deactMultiSend bool, authority string, logger log.Logger,
) BaseKeeper {
	keeper := bankkeeper.NewBaseKeeper(cdc, storeService, ak, blockedAddr, authority, logger)
	baseKeeper := BaseKeeper{
		BaseKeeper:     keeper,
		ak:             ak,
		cdc:            cdc,
		storeService:   storeService,
		inactiveAddrs:  map[string]bool{},
		deactMultiSend: deactMultiSend,
		addrCdc:        cdc.InterfaceRegistry().SigningContext().AddressCodec(),
	}

	keeper.BaseSendKeeper.AppendSendRestriction(func(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) (newToAddr sdk.AccAddress, err error) {
		if baseKeeper.IsInactiveAddr(toAddr) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddr)
		}
		return toAddr, nil
	})
	return baseKeeper
}

func (k BaseKeeper) InitializeBankPlus(ctx context.Context) {
	k.loadAllInactiveAddrs(ctx)
}

// AddToInactiveAddr adds the address to `inactiveAddr`.
func (k BaseKeeper) AddToInactiveAddr(ctx context.Context, addr sdk.AccAddress) {
	addrString, err := k.addrCdc.BytesToString(addr)
	if err != nil {
		panic(err)
	}
	if !k.inactiveAddrs[addrString] {
		k.inactiveAddrs[addrString] = true

		k.addToInactiveAddr(ctx, addr)
	}
}

// DeleteFromInactiveAddr removes the address from `inactiveAddr`.
func (k BaseKeeper) DeleteFromInactiveAddr(ctx context.Context, addr sdk.AccAddress) {
	addrString, err := k.addrCdc.BytesToString(addr)
	if err != nil {
		panic(err)
	}
	if k.inactiveAddrs[addrString] {
		delete(k.inactiveAddrs, addrString)

		k.deleteFromInactiveAddr(ctx, addr)
	}
}

// IsInactiveAddr returns if the address is added in inactiveAddr.
func (k BaseKeeper) IsInactiveAddr(addr sdk.AccAddress) bool {
	addrString, err := k.addrCdc.BytesToString(addr)
	if err != nil {
		panic(err)
	}
	return k.inactiveAddrs[addrString]
}

func (k BaseKeeper) InputOutputCoins(ctx context.Context, input types.Input, outputs []types.Output) error {
	if k.deactMultiSend {
		return sdkerrors.ErrNotSupported.Wrap("MultiSend was deactivated")
	}

	for _, out := range outputs {
		if k.inactiveAddrs[out.Address] {
			return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", out.Address)
		}
	}

	return k.BaseSendKeeper.InputOutputCoins(ctx, input, outputs)
}

package bankplus

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModule struct {
	bank.AppModule

	bankKeeper     bankkeeper.Keeper
	legacySubspace exported.Subspace
}

func NewAppModule(cdc codec.Codec, keeper bankkeeper.Keeper, accountKeeper accountkeeper.AccountKeeper, ss exported.Subspace) AppModule {
	return AppModule{
		AppModule:      bank.NewAppModule(cdc, keeper, accountKeeper, ss),
		bankKeeper:     keeper,
		legacySubspace: ss,
	}
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	banktypes.RegisterMsgServer(cfg.MsgServer(), bankkeeper.NewMsgServerImpl(am.bankKeeper))
	banktypes.RegisterQueryServer(cfg.QueryServer(), am.bankKeeper)

}

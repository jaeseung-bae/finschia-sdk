package bankplus

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type AppModule struct {
	bank.AppModule

	bankKeeper bankkeeper.Keeper
}

package querier

import (
	"github.com/line/link/x/token/internal/keeper"
	"github.com/line/link/x/token/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// creates a querier for token REST endpoints
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryPerms:
			return queryAccountPermission(ctx, req, keeper)
		case types.QueryTokens:
			return queryTokens(ctx, req, keeper)
		case types.QueryBalance:
			return queryBalance(ctx, req, keeper)
		case types.QueryMint:
			return queryTotal(ctx, req, keeper, types.QueryMint)
		case types.QueryBurn:
			return queryTotal(ctx, req, keeper, types.QueryBurn)
		case types.QuerySupply:
			return queryTotal(ctx, req, keeper, types.QuerySupply)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown token query endpoint")
		}
	}
}

func queryBalance(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, error) {
	if len(req.Data) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "data is nil")
	}
	var params types.QueryAccAddressContractIDParams
	if err := keeper.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	supply := keeper.GetBalance(ctx, params.ContractID, params.Addr)
	bz, err := keeper.MarshalJSONIndent(supply)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAccountPermission(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, error) {
	if len(req.Data) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "data is nil")
	}
	var params types.QueryAccAddressContractIDParams
	if err := keeper.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	pms := keeper.GetPermissions(ctx, params.ContractID, params.Addr)

	bz, err := keeper.MarshalJSONIndent(pms)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryTokens(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, error) {
	if len(req.Data) == 0 {
		tokens := keeper.GetAllTokens(ctx)

		bz, err := keeper.MarshalJSONIndent(tokens)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}
	var params types.QueryContractIDParams
	if err := keeper.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	token, err := keeper.GetToken(ctx, params.ContractID)
	if err != nil {
		return nil, err
	}

	bz, err2 := keeper.MarshalJSONIndent(token)
	if err2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err2.Error())
	}

	return bz, nil
}

func queryTotal(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper, target string) ([]byte, error) {
	var params types.QueryContractIDParams
	if err := keeper.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	total, err := keeper.GetTotalInt(ctx, params.ContractID, target)
	if err != nil {
		return nil, err
	}

	bz, err2 := keeper.MarshalJSONIndent(total)
	if err2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err2.Error())
	}

	return bz, nil
}

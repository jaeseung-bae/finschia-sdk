package keeper_test

import (
	"context"
	"testing"

	ocproto "github.com/line/ostracon/proto/ostracon/types"

	"github.com/stretchr/testify/suite"

	"github.com/line/lbm-sdk/crypto/keys/secp256k1"
	"github.com/line/lbm-sdk/simapp"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/collection"
	"github.com/line/lbm-sdk/x/collection/keeper"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx         sdk.Context
	goCtx       context.Context
	keeper      keeper.Keeper
	queryServer collection.QueryServer
	msgServer   collection.MsgServer

	vendor   sdk.AccAddress
	operator sdk.AccAddress
	customer sdk.AccAddress
	stranger sdk.AccAddress

	contractID string
	ftClassID  string
	nftClassID string

	balance sdk.Int

	numNFTs int
}

func createRandomAccounts(accNum int) []sdk.AccAddress {
	seenAddresses := make(map[sdk.AccAddress]bool, accNum)
	addresses := make([]sdk.AccAddress, accNum)
	for i := 0; i < accNum; i++ {
		var address sdk.AccAddress
		for {
			pk := secp256k1.GenPrivKey().PubKey()
			address = sdk.BytesToAccAddress(pk.Address())
			if !seenAddresses[address] {
				seenAddresses[address] = true
				break
			}
		}
		addresses[i] = address
	}
	return addresses
}

func (s *KeeperTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)
	s.ctx = app.BaseApp.NewContext(checkTx, ocproto.Header{})
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	s.keeper = app.CollectionKeeper

	s.queryServer = keeper.NewQueryServer(s.keeper)
	s.msgServer = keeper.NewMsgServer(s.keeper)

	addresses := []*sdk.AccAddress{
		&s.vendor,
		&s.operator,
		&s.customer,
		&s.stranger,
	}
	for i, address := range createRandomAccounts(len(addresses)) {
		*addresses[i] = address
	}

	s.balance = sdk.NewInt(1000000)

	// create a contract
	s.contractID = s.keeper.CreateContract(s.ctx, s.vendor, collection.Contract{
		Name: "fox",
	})

	for _, permission := range []collection.Permission{
		collection.PermissionMint,
		collection.PermissionBurn,
	} {
		s.keeper.Grant(s.ctx, s.contractID, s.vendor, s.operator, permission)
	}

	// create a fungible token class
	ftClassID, err := s.keeper.CreateTokenClass(s.ctx, s.contractID, &collection.FTClass{
		Name:     "tibetian fox",
		Mintable: true,
	})
	s.Require().NoError(err)
	s.ftClassID = *ftClassID

	// create a non-fungible token class
	nftClassID, err := s.keeper.CreateTokenClass(s.ctx, s.contractID, &collection.NFTClass{
		Name: "fennec fox",
	})
	s.Require().NoError(err)
	s.nftClassID = *nftClassID

	// mint & burn fts
	for _, to := range []sdk.AccAddress{s.customer, s.operator, s.vendor} {
		tokenID := collection.NewFTID(s.ftClassID)
		amount := collection.NewCoins(collection.NewCoin(tokenID, s.balance))

		err := s.keeper.MintFT(s.ctx, s.contractID, to, amount)
		s.Require().NoError(err)

		_, err = s.keeper.BurnCoins(s.ctx, s.contractID, to, amount)
		s.Require().NoError(err)
		err = s.keeper.MintFT(s.ctx, s.contractID, to, amount)
		s.Require().NoError(err)
	}

	// mint nfts
	newParams := func(classID string, size int) []collection.MintNFTParam {
		res := make([]collection.MintNFTParam, size)
		for i := range res {
			res[i] = collection.MintNFTParam{
				TokenType: s.nftClassID,
			}
		}
		return res
	}
	// 1 for the successful attach, 2 for the failure
	remainders := 1 + 2
	s.numNFTs = collection.DefaultDepthLimit + remainders
	for _, to := range []sdk.AccAddress{s.customer, s.operator, s.vendor} {
		tokens, err := s.keeper.MintNFT(s.ctx, s.contractID, to, newParams(s.nftClassID, collection.DefaultDepthLimit))
		s.Require().NoError(err)

		for i := range tokens[1:] {
			r := len(tokens) - 1 - i
			subject := tokens[r].Id
			target := tokens[r-1].Id
			err := s.keeper.Attach(s.ctx, s.contractID, to, subject, target)
			s.Require().NoError(err)
		}

		tokens, err = s.keeper.MintNFT(s.ctx, s.contractID, to, newParams(s.nftClassID, remainders))
		s.Require().NoError(err)

		err = s.keeper.Attach(s.ctx, s.contractID, to, tokens[remainders-1].Id, tokens[remainders-2].Id)
		s.Require().NoError(err)

	}

	// authorize
	err = s.keeper.AuthorizeOperator(s.ctx, s.contractID, s.customer, s.operator)
	s.Require().NoError(err)
	err = s.keeper.AuthorizeOperator(s.ctx, s.contractID, s.customer, s.stranger)
	s.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
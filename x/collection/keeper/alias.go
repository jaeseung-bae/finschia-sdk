package keeper

import (
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/Finschia/finschia-sdk/x/collection"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// iterate through the balances of a contract and perform the provided function
func (k Keeper) iterateContractBalances(ctx sdk.Context, contractID string, fn func(address sdk.AccAddress, balance collection.Coin) (stop bool)) {
	k.iterateBalancesImpl(ctx, balanceKeyPrefixByContractID(contractID), func(_ string, address sdk.AccAddress, balance collection.Coin) (stop bool) {
		return fn(address, balance)
	})
}

func (k Keeper) iterWithPrefix(ctx sdk.Context, prefix []byte) storetypes.Iterator {
	store := k.storeService.OpenKVStore(ctx)
	adapter := runtime.KVStoreAdapter(store)
	iterator := storetypes.KVStorePrefixIterator(adapter, prefix)
	return iterator
}

func (k Keeper) iterateBalancesImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, address sdk.AccAddress, balance collection.Coin) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, address, tokenID := splitBalanceKey(iterator.Key())

		var amount math.Int
		if err := amount.Unmarshal(iterator.Value()); err != nil {
			panic(err)
		}
		balance := collection.NewCoin(tokenID, amount)

		stop := fn(contractID, address, balance)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContracts(ctx sdk.Context, fn func(contract collection.Contract) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, contractKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var contract collection.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)

		stop := fn(contract)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractClasses(ctx sdk.Context, contractID string, fn func(class collection.TokenClass) (stop bool)) {
	k.iterateClassesImpl(ctx, classKeyPrefixByContractID(contractID), fn)
}

// iterate through the classes and perform the provided function
func (k Keeper) iterateClassesImpl(ctx sdk.Context, prefix []byte, fn func(class collection.TokenClass) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var class collection.TokenClass
		if err := k.cdc.UnmarshalInterface(iterator.Value(), &class); err != nil {
			panic(err)
		}

		stop := fn(class)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractGrants(ctx sdk.Context, contractID string, fn func(grant collection.Grant) (stop bool)) {
	k.iterateGrantsImpl(ctx, grantKeyPrefixByContractID(contractID), func(_ string, grant collection.Grant) (stop bool) {
		return fn(grant)
	})
}

func (k Keeper) iterateGrantsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, grant collection.Grant) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, grantee, permission := splitGrantKey(iterator.Key())
		grant := collection.Grant{
			Grantee:    grantee.String(),
			Permission: permission,
		}

		stop := fn(contractID, grant)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractAuthorizations(ctx sdk.Context, contractID string, fn func(authorization collection.Authorization) (stop bool)) {
	k.iterateAuthorizationsImpl(ctx, authorizationKeyPrefixByContractID(contractID), func(_ string, authorization collection.Authorization) (stop bool) {
		return fn(authorization)
	})
}

func (k Keeper) iterateAuthorizationsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, authorization collection.Authorization) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, operator, holder := splitAuthorizationKey(iterator.Key())
		authorization := collection.Authorization{
			Holder:   holder.String(),
			Operator: operator.String(),
		}

		stop := fn(contractID, authorization)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractNFTs(ctx sdk.Context, contractID string, fn func(nft collection.NFT) (stop bool)) {
	k.iterateNFTsImpl(ctx, nftKeyPrefixByContractID(contractID), func(_ string, nft collection.NFT) (stop bool) {
		return fn(nft)
	})
}

func (k Keeper) iterateNFTsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, NFT collection.NFT) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, _ := splitNFTKey(iterator.Key())

		var nft collection.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &nft)

		if fn(contractID, nft) {
			break
		}
	}
}

func (k Keeper) iterateContractParents(ctx sdk.Context, contractID string, fn func(tokenID, parentID string) (stop bool)) {
	k.iterateParentsImpl(ctx, parentKeyPrefixByContractID(contractID), func(_ string, tokenID, parentID string) (stop bool) {
		return fn(tokenID, parentID)
	})
}

func (k Keeper) iterateParentsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, tokenID, parentID string) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, tokenID := splitParentKey(iterator.Key())

		var parentID gogotypes.StringValue
		k.cdc.MustUnmarshal(iterator.Value(), &parentID)

		if fn(contractID, tokenID, parentID.Value) {
			break
		}
	}
}

func (k Keeper) iterateChildrenImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, tokenID, childID string) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, tokenID, childID := splitChildKey(iterator.Key())
		if fn(contractID, tokenID, childID) {
			break
		}
	}
}

func (k Keeper) iterateStatisticsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, classID string, amount math.Int) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var amount math.Int
		if err := amount.Unmarshal(iterator.Value()); err != nil {
			panic(err)
		}

		keyPrefix := prefix[:1]
		contractID, classID := splitStatisticKey(keyPrefix, iterator.Key())

		stop := fn(contractID, classID, amount)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractSupplies(ctx sdk.Context, contractID string, fn func(classID string, amount math.Int) (stop bool)) {
	k.iterateStatisticsImpl(ctx, statisticKeyPrefixByContractID(supplyKeyPrefix, contractID), func(_ string, classID string, amount math.Int) (stop bool) {
		return fn(classID, amount)
	})
}

func (k Keeper) iterateContractBurnts(ctx sdk.Context, contractID string, fn func(classID string, amount math.Int) (stop bool)) {
	k.iterateStatisticsImpl(ctx, statisticKeyPrefixByContractID(burntKeyPrefix, contractID), func(_ string, classID string, amount math.Int) (stop bool) {
		return fn(classID, amount)
	})
}

// iterate through the next token class ids and perform the provided function
func (k Keeper) iterateNextTokenClassIDs(ctx sdk.Context, fn func(class collection.NextClassIDs) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, nextClassIDKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var class collection.NextClassIDs
		k.cdc.MustUnmarshal(iterator.Value(), &class)

		stop := fn(class)
		if stop {
			break
		}
	}
}

func (k Keeper) iterateContractNextTokenIDs(ctx sdk.Context, contractID string, fn func(nextID collection.NextTokenID) (stop bool)) {
	k.iterateNextTokenIDsImpl(ctx, nextTokenIDKeyPrefixByContractID(contractID), func(_ string, nextID collection.NextTokenID) (stop bool) {
		return fn(nextID)
	})
}

// iterate through the next (non-fungible) token ids and perform the provided function
func (k Keeper) iterateNextTokenIDsImpl(ctx sdk.Context, prefix []byte, fn func(contractID string, nextID collection.NextTokenID) (stop bool)) {
	iterator := k.iterWithPrefix(ctx, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractID, classID := splitNextTokenIDKey(iterator.Key())

		var id math.Uint
		if err := id.Unmarshal(iterator.Value()); err != nil {
			panic(err)
		}

		nextID := collection.NextTokenID{
			ClassId: classID,
			Id:      id,
		}

		stop := fn(contractID, nextID)
		if stop {
			break
		}
	}
}

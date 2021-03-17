package store

import (
	dbm "github.com/line/tm-db/v2"

	"github.com/line/lbm-sdk/v2/store/cache"
	"github.com/line/lbm-sdk/v2/store/rootmulti"
	"github.com/line/lbm-sdk/v2/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}

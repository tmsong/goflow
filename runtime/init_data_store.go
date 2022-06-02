package runtime

import (
	redisDataStore "github.com/tmsong/goflow/core/redis-datastore"
	"github.com/tmsong/goflow/core/sdk"
)

func initDataStore(redisURI string) (dataStore sdk.DataStore, err error) {
	dataStore, err = redisDataStore.GetRedisDataStore(redisURI)
	return dataStore, err
}

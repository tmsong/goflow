package runtime

import (
	redisStateStore "github.com/tmsong/goflow/core/redis-statestore"
	"github.com/tmsong/goflow/core/sdk"
)

func initStateStore(redisURI string) (stateStore sdk.StateStore, err error) {
	stateStore, err = redisStateStore.GetRedisStateStore(redisURI)
	return stateStore, err
}

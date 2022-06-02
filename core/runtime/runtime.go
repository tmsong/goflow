package runtime

import (
	"github.com/tmsong/goflow/core/sdk/executor"
)

type Runtime interface {
	Init() error
	CreateExecutor(*Request) (executor.Executor, error)
}

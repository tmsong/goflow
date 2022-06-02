package runtime

import (
	"github.com/tmsong/goflow/core/runtime/controller/handler"
	"net/http"

	"github.com/tmsong/goflow/core/runtime"

	"github.com/julienschmidt/httprouter"
)

func router(runtime runtime.Runtime) http.Handler {
	router := httprouter.New()
	router.POST("/:flowName", newRequestHandlerWrapper(runtime, handler.ExecuteFlowHandler))
	router.GET("/:flowName", newRequestHandlerWrapper(runtime, handler.ExecuteFlowHandler))
	return router
}

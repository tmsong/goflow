package runtime

import (
	"fmt"
	"io/ioutil"
	"net/http"

	runtimepkg "github.com/tmsong/goflow/core/runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/tmsong/goflow/core/sdk/executor"
)

func newRequestHandlerWrapper(runtime runtimepkg.Runtime, handler func(*runtimepkg.Response, *runtimepkg.Request, executor.Executor) error) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		flowName := params.ByName("flowName")

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			handleError(w, fmt.Sprintf("failed to execute request, "+err.Error()))
			return
		}

		reqParams := make(map[string][]string)
		for _, param := range params {
			reqParams[param.Key] = []string{param.Value}
		}

		for key, values := range req.URL.Query() {
			reqParams[key] = values
		}

		response := &runtimepkg.Response{}
		response.RequestID = id
		response.Header = make(map[string][]string)
		request := &runtimepkg.Request{
			Body:      body,
			Header:    req.Header,
			FlowName:  flowName,
			RequestID: id,
			Query:     reqParams,
			RawQuery:  req.URL.RawQuery,
		}

		ex, err := runtime.CreateExecutor(request)
		if err != nil {
			handleError(w, fmt.Sprintf("failed to execute request, "+err.Error()))
			return
		}

		err = handler(response, request, ex)
		if err != nil {
			handleError(w, fmt.Sprintf("request failed to be processed, "+err.Error()))
			return
		}

		headers := w.Header()
		for key, values := range response.Header {
			headers[key] = values
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response.Body)
	}
}

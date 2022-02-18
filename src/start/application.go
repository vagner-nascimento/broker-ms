package start

import "broker/src/presentation/http"

func StartApplication() *chan error {
	return http.StartHttpServer()
}

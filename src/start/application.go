package start

import "broker/src/presentation/http"

func StartApplication() *chan error {
	errs := make(chan error)
	http.StartHttpServer(&errs)

	return &errs
}

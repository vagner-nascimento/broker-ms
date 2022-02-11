package http

import (
	"errors"
	"fmt"
	netHttp "net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

func StartHttpServer(ch *chan error) {
	router := chi.NewRouter()

	router.Use(getMiddlewareList()...)
	router.Route("/", func(r chi.Router) {
		r.Mount("/publish", getPublishRoutes())
	})

	walkThroughRoutes := func(method string, route string, handler netHttp.Handler, middleware ...func(netHttp.Handler) netHttp.Handler) error {
		fmt.Println(fmt.Sprintf("%s %s", method, route))
		return nil
	}

	if err := chi.Walk(router, walkThroughRoutes); err != nil {
		fmt.Println("error on walkThroughRoutes", err)
		*ch <- errors.New("error on try to start http server")
	} else {
		port := 3000

		if envPort := os.Getenv("LISTEN_PORT"); len(envPort) > 0 {
			port, _ = strconv.Atoi(envPort)
		}

		fmt.Println(fmt.Sprintf("http server connected on port: %d", port))

		*ch <- netHttp.ListenAndServe(fmt.Sprintf(":%d", port), router)
	}
}

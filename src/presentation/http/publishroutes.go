package http

import (
	"broker/src/app"
	"broker/src/model"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	netHttp "net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

func getPublishRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", postPublish)

	return router
}

func postPublish(w netHttp.ResponseWriter, r *netHttp.Request) {
	if ps, err := parseAndValidatePublishes(r); err == nil {
		app.NewPublishAdapter()
		pubAdp := app.NewPublishAdapter()

		if err := pubAdp.SaveAll(ps); err == nil {
			fmt.Println("postPublish data saved")
			w.WriteHeader(netHttp.StatusAccepted)
		} else {
			fmt.Println("postPublish error on save data", err)
			writeErrorResponse(w, err)
		}
	} else {
		fmt.Println("postPublish parse errors", err)
		writeErrorResponse(w, err)
	}
}

func parseAndValidatePublishes(r *netHttp.Request) (ps []model.Publish, err error) {
	if ps, err = parsePublishesFromBody(r.Body); err == nil {
		var vErrs []validator.ValidationErrors

		for _, p := range ps {
			if vErr := p.Validate(); vErr != nil {
				fmt.Println("parseAndValidatePublishes validation errors: ", vErr)
				vErrs = append(vErrs, vErr)
			}
		}

		if len(vErrs) > 0 {
			err = errors.New("parseAndValidatePublishes some body itens are invalid")
			ps = nil
		}
	}

	return ps, err
}

func parsePublishesFromBody(b io.ReadCloser) (ps []model.Publish, err error) {
	defer b.Close()

	var bys []byte
	if bys, err = ioutil.ReadAll(b); err == nil {
		fmt.Println("parsePublishesFromBody received body", string(bys))

		ps, err = model.NewPublishesFromJsonBytes(bys)
	}

	return ps, err
}

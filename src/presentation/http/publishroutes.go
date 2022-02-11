package http

import (
	"broker/src/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
		fmt.Println("parsed data", ps)
		w.WriteHeader(netHttp.StatusAccepted)
	} else {
		fmt.Println("parse errors", err)

		// TODO: improve bad request error with validation datails
		errRes := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		errJs, _ := json.Marshal(errRes)

		w.WriteHeader(netHttp.StatusBadRequest)
		w.Write(errJs)
	}
}

func parseAndValidatePublishes(r *http.Request) (ps []model.Publish, err error) {
	if ps, err = parsePublishesFromBody(r.Body); err == nil {
		var vErrs []validator.ValidationErrors

		for _, p := range ps {
			if vErr := p.Validate(); vErr != nil {
				fmt.Println("validation errors: ", vErr)
				vErrs = append(vErrs, vErr)
			}
		}

		if len(vErrs) > 0 {
			err = errors.New("some body itens are invalid")
			ps = nil
		}
	}

	return ps, err
}

func parsePublishesFromBody(b io.ReadCloser) (ps []model.Publish, err error) {
	defer b.Close()

	var bys []byte
	if bys, err = ioutil.ReadAll(b); err == nil {
		ps, err = model.NewPublishesFromJsonBytes(bys)
	}
	return ps, err
}

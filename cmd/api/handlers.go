package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/smetanamolokovich/mustafar_task/pkg/data"
	"github.com/smetanamolokovich/mustafar_task/pkg/kvstore"
	"github.com/smetanamolokovich/mustafar_task/pkg/validator"
)

func (app *application) getValueHandler(w http.ResponseWriter, r *http.Request) {
	key, err := app.readKeyParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	value, err := app.store.GetValue(key)

	if err != nil {
		if errors.Is(err, kvstore.ErrNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}

		return
	}
	decVal, err := app.decodeBase64(string(value))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/value/%s", key))

	err = app.writeJSON(w, http.StatusOK, envelope{"value": string(decVal)}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) setValue(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Key     string    `json:"key"`
		Value   []byte    `json:"value"`
		Expires time.Time `json:"expires"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	d := &data.Data{
		Key:     input.Key,
		Value:   input.Value,
		Expires: input.Expires,
	}

	d.Value = []byte(app.encodeBase64(d.Value))

	v := validator.New()

	if data.ValidateData(v, d); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.store.SetValue(d.Key, d.Value, d.Expires)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": input}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

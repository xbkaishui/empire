package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/empire/empire"
)

type GetApps struct {
	Empire
}

func (h *GetApps) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	apps, err := h.AppsAll()
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	return Encode(w, apps)
}

type DeleteApp struct {
	Empire
}

func (h *DeleteApp) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	name := empire.AppName(vars["app"])

	a, err := h.AppsFind(name)
	if err != nil {
		return err
	}

	if a == nil {
		return ErrNotFound
	}

	if err := h.AppsDestroy(a); err != nil {
		return err
	}

	return NoContent(w)
}

type PostAppsForm struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
}

type PostApps struct {
	Empire
}

func (h *PostApps) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var form PostAppsForm

	if err := Decode(r, &form); err != nil {
		return err
	}

	app, err := empire.NewApp(empire.AppName(form.Name), empire.Repo(form.Repo))
	if err != nil {
		return ErrBadRequest
	}

	a, err := h.AppsCreate(app)
	if err != nil {
		return err
	}

	w.WriteHeader(201)
	return Encode(w, a)
}
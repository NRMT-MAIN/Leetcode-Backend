package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(submissionRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	submissionRouter.Register(chiRouter)
	return chiRouter
}
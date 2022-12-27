package http

import (
	"github.com/go-chi/chi/v5"
	"go_cyfra_homework_2/internal/action"
	"go_cyfra_homework_2/internal/service/store"
	"go_cyfra_homework_2/pkg/response"
	"net/http"
)

type ServerImplementation struct {
	db *store.Store
}

func New() ServerImplementation {
	return ServerImplementation{}
}

func (s *ServerImplementation) BuildRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		result, exists := action.GetUser(r)

		if !exists {
			response.Make(w, "User doesn't exist", nil)
		} else {
			response.Make(w, result, nil)
		}
	})

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		result := action.CreateUser(r)

		response.Make(w, result, nil)
	})

	router.Put("/user", func(w http.ResponseWriter, r *http.Request) {
		result := action.ChangeUser(r)
		response.Make(w, result, nil)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		result := action.Login()
		response.Make(w, result, nil)
	})

	return router
}

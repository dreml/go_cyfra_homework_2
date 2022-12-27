package action

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go_cyfra_homework_2/internal/model"
	"go_cyfra_homework_2/internal/service/store"
	"net/http"
)

func GetUser(r *http.Request) (model.User, bool){
	userId, _ := uuid.Parse(chi.URLParam(r, "userId"))

	return store.Storage.User(userId)
}

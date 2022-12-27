package action

import (
	"encoding/json"
	"github.com/google/uuid"
	"go_cyfra_homework_2/internal/model"
	"go_cyfra_homework_2/internal/service/store"
	"net/http"
)

func CreateUser(r *http.Request) uuid.UUID {
	var userData model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}

	userId := store.Storage.CreateUser(userData)

	return userId
}

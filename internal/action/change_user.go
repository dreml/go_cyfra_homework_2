package action

import (
	"encoding/json"
	"go_cyfra_homework_2/internal/model"
	"go_cyfra_homework_2/internal/service/store"
	"net/http"
)

func ChangeUser(r *http.Request) model.User {
	var userData model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}

	isChanged := store.Storage.ChangeUser(userData)

	var user model.User
	if isChanged {
		user, _ = store.Storage.User(userData.Id)
	}

	return user
}

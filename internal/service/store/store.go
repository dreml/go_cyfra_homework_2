package store

import (
	"github.com/google/uuid"
	"go_cyfra_homework_2/internal/model"
	"sync"
)

type Store struct {
	users map[string]model.User
	mx    sync.Mutex
}

var Storage *Store

func New() {
	Storage = &Store{users: map[string]model.User{}}
}

func (s *Store) Wipe() {
	s.users = map[string]model.User{}
}

func (s *Store) User(uuid uuid.UUID) (model.User, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()

	user, ok := s.users[uuid.String()]
	if !ok {
		return model.User{}, false
	}

	return user, true
}

func (s *Store) CreateUser(user model.User) uuid.UUID {
	s.mx.Lock()
	defer s.mx.Unlock()

	userId := uuid.New()
	user.Id = userId

	s.users[userId.String()] = user

	return userId
}
func (s *Store) ChangeUser(user model.User) bool {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.users[user.Id.String()]; !ok {
		return false
	}

	s.users[user.Id.String()] = user

	return true
}
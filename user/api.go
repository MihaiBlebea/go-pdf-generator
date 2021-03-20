package user

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Store struct {
	users  []User
	logger Logger
}

func NewStore(logger Logger) *Store {
	return &Store{logger: logger}
}

func (s *Store) GenerateUsers(count int) {
	s.logger.Info("Starting generating users")
	for i := 0; i < count; i++ {
		s.users = append(s.users, User{
			ID:        uuid.NewV4().String(),
			FirstName: "John",
			LastName:  "Doe",
		})
	}
	s.logger.Info("Completed generating users")
}

func (s *Store) Users() []User {
	return s.users
}

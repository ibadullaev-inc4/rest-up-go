package store_test

import (
	"rest-up-go/internal/app/model"
	"testing"

	"rest-up-go/internal/app/store"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teaddown := store.TestStore(t, databaseURL)
	defer teaddown("users")

	user, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teaddown := store.TestStore(t, databaseURL)
	defer teaddown("users")

	email := "non-existent-user@email.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	user, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

}

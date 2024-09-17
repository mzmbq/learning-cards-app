package teststore_test

import (
	"testing"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Find(t *testing.T) {
	repo := teststore.NewUserRepo(nil)

	u := model.User{
		ID:    123,
		Email: "alice@bob.com",
	}

	unknownUser := model.User{
		ID: 456,
	}

	err := repo.Create(&u)
	assert.Nil(t, err)

	foundUser, err := repo.Find(u.ID)
	assert.Nil(t, err)
	assert.Equal(t, foundUser.ID, u.ID)

	_, err = repo.Find(unknownUser.ID)
	assert.NotNil(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	repo := teststore.NewUserRepo(nil)

	u := model.User{
		ID:    123,
		Email: "alice@bob.com",
	}

	unknownUser := model.User{
		Email: "who@am.i",
	}

	err := repo.Create(&u)
	assert.Nil(t, err)

	foundUser, err := repo.FindByEmail(u.Email)
	assert.Nil(t, err)
	assert.Equal(t, foundUser.ID, u.ID)

	foundUser, err = repo.FindByEmail(unknownUser.Email)
	assert.NotNil(t, err)
	assert.Nil(t, foundUser)
}

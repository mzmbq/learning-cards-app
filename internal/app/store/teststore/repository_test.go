package teststore_test

import (
	"testing"

	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store/teststore"
)

func TestUserRepository_Find(t *testing.T) {
	repo := teststore.NewUserRepo(nil)

	u := model.User{
		ID:    123,
		Email: "alice@bob.com",
	}

	unknownUser := model.User{
		ID:    456,
		Email: "who@am.i",
	}

	repo.Create(&u)

	foundUser, err := repo.Find(u.ID)
	if err != nil {
		t.Fatal()
	}

	if foundUser.ID != u.ID {
		t.Fatal()
	}

	_, err = repo.Find(unknownUser.ID)
	if err == nil {
		t.Fatal()
	}

}

func TestUserRepository_FindByEmail(t *testing.T) {
	repo := teststore.NewUserRepo(nil)

	u := model.User{
		ID:    123,
		Email: "alice@bob.com",
	}

	unknownUser := model.User{
		ID:    456,
		Email: "who@am.i",
	}

	repo.Create(&u)

	foundUser, err := repo.FindByEmail(u.Email)
	if err != nil {
		t.Fatal()
	}

	if foundUser.ID != u.ID {
		t.Fatal()
	}

	foundUser, err = repo.FindByEmail(unknownUser.Email)
	if err == nil {
		t.Fatal(foundUser)
	}
}

package teststore

import (
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func NewUserRepo(s *Store) *UserRepository {
	return &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
}

func (r *UserRepository) Create(u *model.User) error {
	u.BeforeCreate()

	u.ID = len(r.users) + 1
	r.users[u.Email] = u

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

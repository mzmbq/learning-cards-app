package teststore

import (
	"github.com/mzmbq/learning-cards-app/backend/internal/app/model"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func newUserRepo(s *Store) *UserRepository {
	return &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
}

func (r *UserRepository) Create(u *model.User) error {
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

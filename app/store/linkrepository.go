package store

import "go-link-shorterer/app/model"

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create (l *model.Link) (*model.Link, error){
	if err := r.store.db.QueryRow("asdasd RETURNING id"); err != nil {
		return nil,nil
	}

	return nil, nil
}

func (r *LinkRepository) FindExistUrl(url string) (*model.Link, error) {
	return nil, nil
}

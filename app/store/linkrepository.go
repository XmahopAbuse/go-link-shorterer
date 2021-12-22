package store

import "go-link-shorterer/app/model"

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create (l *model.Link) (*model.Link, error){
	if err := r.store.db.QueryRow(
		"INSERT into links(url, shortUrl) VALUES ($1, $2)",
			l.Url, l.ShortUrl);
	err != nil {
		return nil,nil
	}

	return nil, nil
}

func (r *LinkRepository) FindExistUrl(url string) (*model.Link, error) {
	return nil, nil
}

package store

import "go-link-shorterer/app/model"

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create(l *model.Link) (*model.Link, error) {

	if err := r.store.db.QueryRow("INSERT INTO links (short_url, url) values ($1, $2) RETURNING url, short_url", l.ShortUrl, l.Url).Scan(&l.Url, &l.ShortUrl); err != nil {
		return nil, err
	}

	return l, nil
}

func (r *LinkRepository) FindExistUrl(url string) (*model.Link, error) {
	l := &model.Link{}

	if err := r.store.db.QueryRow("SELECT * from links where url = $1", url).Scan(&l.Url, &l.ShortUrl); err != nil {
		return nil, err
	}

	return l, nil
}

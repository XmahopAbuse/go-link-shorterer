package model

import "go-link-shorterer/app/utils"

type Link struct {
	Url      string
	ShortUrl string
}

func (l *Link) BeforeCreate() error {
	if len(l.Url) > 0 {
		l.ShortUrl = utils.GenerateShortUrl()
	}
	return nil
}

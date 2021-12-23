package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go-link-shorterer/app/utils"
)

type Link struct {
	Url      string
	ShortUrl string
}

func (l *Link) Validate() error {
	return validation.ValidateStruct(l, validation.Field(&l.Url, validation.Required, is.URL))
}

func (l *Link) GenerateShortUrl() {
	if len(l.Url) > 0 {
		l.ShortUrl = utils.GenerateShortUrl()
	}
}

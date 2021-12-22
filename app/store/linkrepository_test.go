package store_test

import (
	"github.com/stretchr/testify/assert"
	"go-link-shorterer/app/model"
	"go-link-shorterer/app/store"
	"go-link-shorterer/app/utils"
	"testing"
)

func TestLinkRepository_Create(t *testing.T) {
	s, _ := store.TestStore(t, databaseUrl)
	l, err := s.Link().Create(&model.Link{
		Url:      "example.com",
		ShortUrl: utils.GenerateShortUrl(),
	})
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

func TestLinkRepository_FindExistUrl(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("links")
	url := "example.com"

	l, err := s.Link().FindExistUrl(url)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

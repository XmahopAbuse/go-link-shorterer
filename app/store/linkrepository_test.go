package store_test

import (
	"github.com/stretchr/testify/assert"
	"go-link-shorterer/app/model"
	"go-link-shorterer/app/store"
	"testing"
)

func TestLinkRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t , databaseUrl)
	defer teardown("links")

	l, err := s.Link().Create(&model.Link{
		Url: "https://google.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, l)
}

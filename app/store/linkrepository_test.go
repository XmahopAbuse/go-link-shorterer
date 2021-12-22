package store_test

import (
	"github.com/stretchr/testify/assert"
	"go-link-shorterer/app/model"
	"go-link-shorterer/app/store"
	"testing"
)

func TestLinkRepository_Create(t *testing.T) {
	s, _ := store.TestStore(t, databaseUrl)
	l, err := s.Link().Create(model.TestLink(t))
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

func TestLinkRepository_FindExistUrl(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("links")
	l := model.TestLink(t)
	l, err := s.Link().FindExistUrl(l.Url)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

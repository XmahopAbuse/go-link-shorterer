package model_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-link-shorterer/app/model"
	"testing"
)

func TestLink_GenerateShortUrl(t *testing.T) {
	l := model.TestLink(t)
	assert.NotEmpty(t, l.ShortUrl)
	fmt.Println(l)
}

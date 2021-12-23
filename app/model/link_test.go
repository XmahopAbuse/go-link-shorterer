package model_test

import (
	"github.com/stretchr/testify/assert"
	"go-link-shorterer/app/model"
	"testing"
)

func TestLink_GenerateShortUrl(t *testing.T) {
	l := model.TestLink(t)
	l.GenerateShortUrl()
	assert.NotEmpty(t, l.ShortUrl)
}

func TestLink_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		l       *model.Link
		isValid bool
	}{
		{
			name:    "valid",
			l:       model.TestLink(t),
			isValid: true,
		},
		{
			name: "empty url",
			l: &model.Link{
				Url: "",
			},
			isValid: false,
		},
		{
			name: "invalid url",
			l: &model.Link{
				Url: "asdasdas",
			},
			isValid: false,
		},
		{
			name: "https url",
			l: &model.Link{
				Url: "https://google.com",
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.l.Validate())
			} else {
				assert.Error(t, tc.l.Validate())
			}
		})
	}
}

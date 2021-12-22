package model

import "testing"

// Нужен. чтобы каждый раз возвращаться пользователя с валидным набором данных
func TestLink(t *testing.T) *Link {
	return &Link{
		Url: "http://google.com",
	}
}

package main

import (
	"math/rand"
	"strings"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMOPQRSTUVWXYZ123456789_")
var shortLinkLength int = 10

func generateShortUrl() string {
	rand.Seed(time.Now().UnixNano())
	var shortUrl strings.Builder
	for i := 1; i <= shortLinkLength ; i++ {
		shortUrl.WriteRune(symbols[rand.Intn(len(symbols))])
	}

	return shortUrl.String()

}

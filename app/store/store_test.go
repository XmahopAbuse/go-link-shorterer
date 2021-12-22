package store_test

import (
	"os"
	"testing"
)

const databaseUrl string = "host=localhost dbname=link-shorterer-test sslmode=disable password=yvxn6akk user=xmahop"

func TestMain(m *testing.M) {

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=link-shorterer-test sslmode=disable password=yvxn6akk user=xmahop"
	}

	os.Exit(m.Run())
}

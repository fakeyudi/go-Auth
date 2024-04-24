package auth

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

const (
	key    = "someString"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gcpCLientId := os.Getenv("GCP_CLIENT_ID")
	gcpClientSecret := os.Getenv("GCP_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(google.New(gcpCLientId, gcpClientSecret, "https://localhost:8080/auth/google/callback"))
}

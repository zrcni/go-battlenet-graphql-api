package battlenet

import (
	"bytes"
	"context"
	base64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const battleNetOauthURL = "https://eu.battle.net/oauth/token"

// TIL struct fields need to be exported for json-package to see their value!
type bnetOauthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Auth is used to authenticate with bnet API and persist the token
type Auth struct {
	token string
}

func generateBasicAuthHeader() string {
	battleNetClientID := os.Getenv("BATTLE_NET_CLIENT_ID")
	battleNetClientSecret := os.Getenv("BATTLE_NET_CLIENT_SECRET")

	usernameAndPassword := fmt.Sprintf("%s:%s", battleNetClientID, battleNetClientSecret)
	basicAuth := base64.StdEncoding.EncodeToString([]byte(usernameAndPassword))

	return fmt.Sprintf("Basic %s", basicAuth)
}

// Authenticate with bnet API
func (auth *Auth) Authenticate() error {
	grantType := "grant_type=client_credentials"
	postBody := bytes.NewBuffer([]byte(grantType))

	req, err := http.NewRequest("POST", battleNetOauthURL, postBody)

	if err != nil {
		log.Println(err)
		return nil
	}

	basicAuth := generateBasicAuthHeader()

	req.Header.Add("Authorization", basicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	var oauthResponse bnetOauthResponse
	json.NewDecoder(res.Body).Decode(&oauthResponse)

	if oauthResponse.AccessToken != "" {
		auth.token = oauthResponse.AccessToken
		log.Print("Successfully authenticated with bnet!")
		return nil
	}

	errorMessage := "Could not authenticate with bnet!"
	log.Print(errorMessage)
	return errors.New(errorMessage)
}

// GetToken returns the token
func (auth *Auth) GetToken() string {
	return auth.token
}

// GetAuthFromContext extracts bnet.Auth from context
func GetAuthFromContext(ctx context.Context) *Auth {
	raw, ok := ctx.Value(bnetAuthCtxKey).(*Auth)

	if !ok {
		payload := handler.GetInitPayload(ctx)
		if payload == nil {
			return nil
		}
	}

	return raw
}

var bnetAuthCtxKey = &contextKey{"bnet_auth"}

type contextKey struct {
	name string
}

// Middleware adds bnet.Auth to the context
func Middleware(bnetAuth *Auth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// return new context with previous context values and added value
			ctx := context.WithValue(r.Context(), bnetAuthCtxKey, bnetAuth)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

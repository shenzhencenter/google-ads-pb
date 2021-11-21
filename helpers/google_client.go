package helpers

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"time"
)

func NewToken(ctx context.Context, access string, tokenType string, refresh string, expiry time.Time) (*oauth2.Token, error) {
	cfg := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
	}
	token := oauth2.Token{
		AccessToken:  access,
		TokenType:    tokenType,
		RefreshToken: refresh,
		Expiry:       expiry,
	}

	client := cfg.Client(ctx, &token)
	if !token.Valid() {
		return RefreshAndSaveToken(client)
	} else {
		return &token, nil
	}
}

func RefreshAndSaveToken(client *http.Client) (*oauth2.Token, error) {

	transport, ok := client.Transport.(*oauth2.Transport)
	if !ok {
		return nil, errors.New("client transport is not oauth2.Transport")
	}

	token, err := transport.Source.Token()
	if err != nil {
		return nil, err
	}

	// You may want to save token to database
	//db.Update(map[string]string{
	//	"AccessToken":  token.AccessToken,
	//	"RefreshToken": token.RefreshToken,
	//	"ExpiresAt":    token.Expiry.Format(time.RFC3339),
	//})

	return token, nil
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func main() {
	var clientID string
	var clientSecret string

	flag.StringVar(&clientID, "client-id", "", "oidc client id")
	flag.StringVar(&clientSecret, "client-secret", "", "oidc client secret")

	flag.Parse()
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5556/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		fmt.Println(rawIDToken)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
		kubectlCmd := fmt.Sprintf(`\n\nkubectl config set-credentials %v \
		--auth-provider=oidc \
		--auth-provider-arg=idp-issuer-url=( issuer url ) \
		--auth-provider-arg=client-id=( your client id ) \
		--auth-provider-arg=client-secret=( your client secret ) \
		--auth-provider-arg=refresh-token=( your refresh token ) \
		--auth-provider-arg=idp-certificate-authority=( path to your ca certificate ) \
		--auth-provider-arg=id-token=( your id_token )`, "USERNAME")
		w.Write([]byte(kubectlCmd))
	})

	log.Printf("listening on http://%s/", "127.0.0.1:5556")
	log.Fatal(http.ListenAndServe("127.0.0.1:5556", nil))
}

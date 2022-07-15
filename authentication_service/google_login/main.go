package google_login

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
)

const (
	ClientId          = "991327617354-070psq66mleedhlg5bbmalch1esap5bo.apps.googleusercontent.com"
	ClientSecret      = "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"
	oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type GoogleResponse struct {
	Email string `json:"email"`
}

var oAuthConfig = &oauth2.Config{
	ClientID:     ClientId,
	ClientSecret: ClientSecret,
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:8080/redirect",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email openid "},
}

func AuthGoogle(w http.ResponseWriter, r *http.Request) {

	url := oAuthConfig.AuthCodeURL("state")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func Redirect(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")
	token, err := oAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return
	}
	ts := oAuthConfig.TokenSource(context.Background(), token)

	client := oauth2.NewClient(context.Background(), ts)

	response, err := client.Get(oauthGoogleUrlAPI)
	if err != nil {
		log.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)

	var user GoogleResponse

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		log.Println("Couldn't decode:", err)
	}

	_, err = fmt.Fprint(w, "Welcome,", user.Email)
	if err != nil {
		log.Println(err)
	}

	//TODO check if the users exists

	//TODO if exists create a jwt token - custom login

}

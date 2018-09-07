package authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	randomState = randomString(15)
)

type OauthSrv struct {
	googleOauthConfig *oauth2.Config
}

type OauthResponse struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PictureURL    string `json:"picture"`
}

func NewOauthSrv() *OauthSrv {
	oauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return &OauthSrv{googleOauthConfig: oauthConfig}
}

func (oauth *OauthSrv) OauthCallback(w http.ResponseWriter, r *http.Request) (*OauthResponse, error) {
	if r.FormValue("state") != randomState {
		return nil, fmt.Errorf("state is not valid")
	}

	token, err := oauth.googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "could not get request")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse response")
	}

	var oauthResp OauthResponse
	err = json.Unmarshal(content, &oauthResp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal oauth response")
	}

	return &oauthResp, nil
}

func (oauth *OauthSrv) LoginGoogleProvider(w http.ResponseWriter, r *http.Request) {
	url := oauth.googleOauthConfig.AuthCodeURL(randomState)
	url = fmt.Sprintf("{\"url\": \"%s\"}", url)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(url))
}

func randomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

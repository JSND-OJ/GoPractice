package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var store = sessions.NewCookieStore([]byte("secret"))

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	CallBackURL = "http://localhost:1333/auth/callback"

	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:     "google client id",
		ClientSecret: "google client secret",
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, _ := template.ParseFiles(name)
	tmpl.Execute(w, data)
}

func RenderMainView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "main.html", nil)
}

func RenderAuthView(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := RandToken()
	session.Values["state"] = state
	session.Save(r, w)
	RenderTemplate(w, "auth.html", GetLoginURL(state))
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	state := session.Values["state"]

	delete(session.Values, "state")
	session.Save(r, w)

	if state != r.FormValue("state") {
		http.Error(w, "Invalid session state", http.StatusUnauthorized)
		return
	}

	token, err := OAuthConf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := OAuthConf.Client(context.Background(), token)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var authUser User
	json.Unmarshal(userInfo, &authUser)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session.Values["user"] = authUser.Email
	session.Values["username"] = authUser.Name
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	mux := pat.New()
	mux.HandleFunc("/", RenderMainView)
	mux.HandleFunc("/auth", RenderAuthView)
	mux.HandleFunc("/auth/callback", Authenticate)

	n := negroni.Classic()
	n.UseHandler(mux)
	log.Fatal(http.ListenAndServe(":3000", n))
}

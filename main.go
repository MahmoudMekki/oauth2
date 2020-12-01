package main

import (
	"html/template"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubConfig = oauth2.Config{
	ClientID:     "74ec53f00031abb0fd01",
	ClientSecret: "61e37b0bc212c12a08187cf050638d1e94741f0e",
	Endpoint:     github.Endpoint,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startOauth)
	http.HandleFunc("/welcome", completeOauth)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "template.gohtml", nil)
}

func startOauth(res http.ResponseWriter, req *http.Request) {
	redirectURL := githubConfig.AuthCodeURL("000")
	http.Redirect(res, req, redirectURL, http.StatusSeeOther)
}

func completeOauth(res http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	state := req.FormValue("state")
	if state != "000" {
		http.Error(res, "Erorr in auth", http.StatusBadRequest)
		return
	}
	token, err := githubConfig.Exchange(req.Context(), code)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	client := githubConfig.Client(req.Context(), token)

}

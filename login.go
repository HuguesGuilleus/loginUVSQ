// GUILLEUS Hugues <ghugues@netc.fr>
// 2019, Guilleus Hugues <ghugues@netc.fr>

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var addresse string = "https://wlc1.reseau.uvsq.fr/login.html"

func main() {
	getAddresse()
	login()
}

func getAddresse() {
	fmt.Println("Récupération de l'adresse ...")
	rep, err := http.Get("http://google.com/")
	if err != nil {
		panic(err)
	}
	loc, err := rep.Location()
	if err != nil {
		panic(err)
	}
	addresse = loc.Query().Get("switch_url")
}

func login() {
	time.AfterFunc(3*1000*1000*1000, func() {
		fmt.Fprintln(os.Stderr, "Temps écoulé (2s)")
		os.Exit(1)
	})
	rep, err := http.PostForm(addresse,
		url.Values{
			"username":      {"LOGIN"},
			"password":      {"PASSW"},
			"buttonClicked": {"4"},
			"redirect_url":  {"detectportal.firefox.com/success.txt"},
			"err_flag":      {"0"},
			"Submit":        {"Valider"},
		})
	if err == nil {
		fmt.Println("Connexion validée")
	} else {
		fmt.Fprintln(os.Stderr, "Connexion échoué")
		fmt.Println("Réponse:", rep)
		fmt.Println("Erreur:", err)
		os.Exit(1)
	}
}

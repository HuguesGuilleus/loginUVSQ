// GUILLEUS Hugues <ghugues@netc.fr>
// 2019, Guilleus Hugues <ghugues@netc.fr>
// BSD 3-Clause License

package main

import (
	"./info"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var addresse string = "https://wlc1.reseau.uvsq.fr/login.html"

func init() {
	log.SetFlags(0)
}

func main() {
	log.Println("loginUVSQ v1.1")
	// getAddresse()
	l,p := info.GetInfo()
	login(l, p)
	info.SaveInfo(l, p)
}

func getAddresse() {
	// time.AfterFunc(2*1000*1000*1000, func() {
	// 	fmt.Fprintln(os.Stderr, "Temps écoulé (2s)")
	// 	os.Exit(1)
	// })
	log.Println("Récupération de l'adresse ...")
	rep, err := http.Get("http://google.com/")
	if err != nil {
		panic(err)
	}
	loc, err := rep.Location()
	if err != nil {
		log.Println("Vous devriez être déjà connecté")
		os.Exit(0)
	}
	addresse = loc.Query().Get("switch_url")
}

func login(l, p string) {
	time.AfterFunc(3*1000*1000*1000, func() {
		fmt.Fprintln(os.Stderr, "Temps écoulé (3s)")
		os.Exit(1)
	})
	log.Println("Envoie des informations")
	rep, err := http.PostForm(addresse,
		url.Values{
			"username":      {l},
			"password":      {p},
			"buttonClicked": {"4"},
			"redirect_url":  {"detectportal.firefox.com/success.txt"},
			"err_flag":      {"0"},
			"Submit":        {"Valider"},
		})
	if err == nil {
		log.Println("Connexion validée")
	} else {
		fmt.Fprintln(os.Stderr, "Connexion échoué")
		log.Println("Réponse:", rep)
		log.Println("Erreur:", err)
		os.Exit(1)
	}
}

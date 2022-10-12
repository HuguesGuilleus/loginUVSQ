// GUILLEUS Hugues <ghugues@netc.fr>
// 2019, Guilleus Hugues <ghugues@netc.fr>
// BSD 3-Clause License

package main

import (
	"crypto/tls"
	"fmt"
	"github.com/HuguesGuilleus/loginUVSQ/info"
	"log"
	"net/http"
	"net/url"
	"time"
)

const MAXTIME = 3

var addresse string = "https://wlc1.reseau.uvsq.fr/login.html"

func init() {
	log.SetFlags(0)
}

func main() {
	log.Println("loginUVSQ v1.5")
	getAddresse()
	l, p := info.GetInfo()
	login(l, p)
	info.SaveInfo(l, p)
	fmt.Print("APPUYEZ SUR ENTRER POUR QUITTER ")
	fmt.Scanf("\n")
}

func getAddresse() {
	deadline := time.AfterFunc(MAXTIME*time.Second, func() {
		log.Fatalf("Temps écoulé (%ds)", MAXTIME)
	})
	defer deadline.Stop()

	log.Println("Récupération de l'adresse ...")
	rep, err := http.Get("http://google.com/")
	if err != nil {
		log.Fatal(err)
	}
	loc, err := rep.Location()
	if err != nil {
		log.Fatal("Pas d'entête Location: Vous devriez être déjà connecté")
	}
	addresse = loc.Query().Get("switch_url")
}

func login(l, p string) {
	log.Println("Envoie des informations")

	client := http.Client{Transport: &http.Transport{
		// We use InsecureSkipVerify because the UVSQ TLS certificate are
		// ofter outdated or wrong configurate.
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: MAXTIME * time.Second,
	}}

	rep, err := client.PostForm(addresse,
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
		log.Println("Réponse:", rep)
		log.Println("Erreur:", err)
		log.Fatal("Connexion échouée")
	}
}

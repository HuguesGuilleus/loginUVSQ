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

func main() {
	time.AfterFunc(3 * 1000*1000*1000,func() {
		fmt.Fprintln(os.Stderr, "Temps écoulé (2s)")
		os.Exit(1)
	})
	rep, err := http.PostForm("https://wlc1.reseau.uvsq.fr/login.html",
		url.Values{
			"username": {"LOGIN"},
			"password": {"PASSW"},
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

// GUILLEUS Hugues <ghugues@netc.fr>
// 2019, Guilleus Hugues <ghugues@netc.fr>
// BSD 3-Clause License

package info

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var mustSaveInfo bool = true

func GetInfo() (l, p string) {
	if len(constLogin) != 0 && len(constPassword) != 0 {
		mustSaveInfo = false
		return constLogin, constPassword
	}
	var err bool
	if l, p, err = getSavedInfo(); err {
		mustSaveInfo = false
		return
	}
	l, p = askInfo()
	return l, p
}

func getSavedInfo() (l, p string, ok bool) {
	return "", "", false
}

func askInfo() (l, p string) {
	fmt.Print("Numéro étudiant: ")
	fmt.Scanf("%s", &l)
	fmt.Print("Mot de passe: ")
	fmt.Scanf("%s", &p)
	return
}

func SaveInfo(l, p string) {
	if mustSaveInfo {
		file := getHome() + ".loginUVSQ.txt"
		log.Println("Save Info in: ", file)
		err := ioutil.WriteFile(file, []byte(l+"\n"+p+"\n"), 0600)
		if err != nil {
			log.Fatal("Save Info", err)
		}
	}
}

// Get HOME envionment variable
func getHome() string {
	val, ok := os.LookupEnv("HOME")
	if ok {
		return val + "/"
	}
	val, ok = os.LookupEnv("home")
	if ok {
		return val + "/"
	}
	return "./"
}

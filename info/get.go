// GUILLEUS Hugues <ghugues@netc.fr>
// 2019, Guilleus Hugues <ghugues@netc.fr>
// BSD 3-Clause License

package info

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	mustSaveInfo = true
	fileName     = ""
)

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	log.Print("read file: ", getFileName())
	file, err := ioutil.ReadFile(getFileName())
	if err != nil {
		panic(err)
	}

	data := bytes.Split(file, []byte{'\n'})
	if len(data) < 2 {
		panic("Bad syntax (login puis mot de passe sur deux lignes)")
	}

	return string(data[0]), string(data[1]), true
}

func askInfo() (l, p string) {
	fmt.Print("Numéro étudiant: ")
	for len(l) == 0 {
		fmt.Scanf("%s", &l)
	}
	fmt.Print("Mot de passe: ")
	for len(p) == 0 {
		fmt.Scanf("%s", &p)
	}
	return
}

func SaveInfo(l, p string) {
	if mustSaveInfo {
		file := getFileName()
		log.Println("Save Info in: ", file)
		err := ioutil.WriteFile(file, []byte(l+"\n"+p+"\n"), 0600)
		if err != nil {
			log.Fatal("Save Info", err)
		}
	}
}

// The file name where are saved the loggin and passwd
func getFileName() string {
	if len(fileName) == 0 {
		return getHome() + ".loginUVSQ.txt"
	} else {
		return fileName
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

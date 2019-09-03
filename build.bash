#!/bin/bash
# GUILLEUS Hugues <ghugues@netc.fr>
# BSD 3-Clause License

# Récupère les information pour la compilation
printf "login UVSQ: "
read login
printf "password UVSQ: "
read passw

# Créé un répertoire pour la compilation
mkdir -p tmp/
cd tmp/
# Injecte le login et le mot de passe dans une copie de login.go
sed -e "s/LOGIN/${login}/" -e "s/PASSW/${passw}/" ../login.go > login.go
# Compile
go build -o ../login login.go

# Netoyage
cd ..
rm -r tmp/
login=""
passw=""

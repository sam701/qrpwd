package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/mdp/qrterminal"
	"github.com/mdp/rsc/qr"
)

func main() {
	fmt.Print("Type your password: ")

	passb, err := gopass.GetPasswd()
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	pass := strings.TrimSpace(string(passb))

	qrterminal.Generate(pass, qr.H, os.Stdout)
}

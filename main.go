package main

import (
	"fmt"
	"log"
	"os"

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
	qrterminal.Generate(string(passb), qr.H, os.Stdout)
}

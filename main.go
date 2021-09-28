package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/mdp/qrterminal"
	"github.com/mdp/rsc/qr"
	"log"
	"os"
)

func trimEscaped(bb []byte) []byte {
	startIx := 0
	for bb[startIx] == 27 {
		startIx += 3
	}
	return bb[startIx:]
}

func main() {
	fmt.Print("Type your password: ")

	passb, err := gopass.GetPasswd()
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	str := string(trimEscaped(passb))
	qrterminal.Generate(str, qr.H, os.Stdout)
}

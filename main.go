package main

import (
	"fmt"
	"regexp"
	"log"

	nere "github.com/hiddn/go_nere/nere"
)

func main() {
	var nas string
	var date_naissance string
	for {
		fmt.Printf("NAS (format : 123456789) : ")
		fmt.Scanln(&nas)
		match, err := regexp.MatchString("^[0-9]{9}$", nas)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			fmt.Println("Le NAS n'est pas valide. Il doit être constitué de 9 chiffres")
		} else {
			break
		}
	}
	for {
		fmt.Printf("Date de naissance (format : AAAA-MM-DD) : ")
		fmt.Scanln(&date_naissance)
		match, err := regexp.MatchString("^[1-2][0-9]{3}-(0[1-9]|1[0-2])-([0-2][0-9]|3[0-1])$", date_naissance)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			fmt.Println("La date de naissance n'est pas valide. Format accepté : AAAA-MM-DD")
		} else {
			break
		}
	}
	nad := nere.ObtenirNAD(nas + date_naissance)
	fmt.Printf("NAD: %s\n", nad)
	appuyersurentrer()
}

func appuyersurentrer() {
	var onsenfou string
	fmt.Print("\nAppuyer sur [entrer] pour continuer...")
	fmt.Scanln(&onsenfou)
}
package main

import (
	"fmt"

	nere "github.com/hiddn/go_nere/nere"
)

func main() {
	var nas string
	var date_naissance string
	fmt.Printf("NAS (format : 123456789) : ")
	fmt.Scanln(&nas)
	fmt.Printf("Date de naissance (format : AAAA-MM-DD) : ")
	fmt.Scanln(&date_naissance)
	nad := nere.ObtenirNAD(nas + date_naissance)
	fmt.Printf("NAD: %s\n", nad)
}

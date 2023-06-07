package main

import (
	"fmt"
	"github.com/pactus-project/pactus/util/bech32m"
	"log"
)

func main() {

	penumbra_addr := "penumbrav2t1fc6gvmyz749cvf7qyz2cgqyw8aq2zf7cm005uvmt4l5dtew5xuvdtuvdrcpwn740xlgx9saeyqtqftwnw57q3vkyd73teckwm9jkwcmwcxml7q7klu9smekthxpa2575urjltu"
	fmt.Println("Penumbra address is:", penumbra_addr)
	hrp, as_b, err := bech32m.DecodeNoLimit(penumbra_addr)
	fmt.Println("Penumbra human-readable prefix:", hrp)
	fmt.Println("Penumbra address decoded to bytes looks like:", as_b)

	// Try via bech32m
	var p2 string
	fmt.Println("Using bech32m nolimit conversion")
	p2, err = bech32m.Encode(hrp, as_b)
	if err != nil {
		log.Fatal("Failed to convert addr from bytes via bech32m")
	}
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Printf("As expected, the two addresses were identical round-trip::\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	}
}

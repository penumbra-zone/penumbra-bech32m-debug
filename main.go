package main

import (
	"fmt"
	"github.com/pactus-project/pactus/util/bech32m"
	"log"
	penumbracrypto "penumbra/address/penumbraprotos/penumbra/core/crypto/v1alpha1"
	"reflect"
)

func bech32m_round_trip() {
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

func via_protos() {
	penumbra_addr := "penumbrav2t1fc6gvmyz749cvf7qyz2cgqyw8aq2zf7cm005uvmt4l5dtew5xuvdtuvdrcpwn740xlgx9saeyqtqftwnw57q3vkyd73teckwm9jkwcmwcxml7q7klu9smekthxpa2575urjltu"
	var a1 penumbracrypto.Address
	var a2 penumbracrypto.Address
	var a3 penumbracrypto.Address
	// Use the new alt_bech32m field supported in the Address proto:
	// https://buf.build/penumbra-zone/penumbra/docs/e5ff7074e14f44328ed2975f01a36f26:penumbra.core.crypto.v1alpha1#penumbra.core.crypto.v1alpha1.Address
	// N.B. This didn't make it into v0.54.1, but it'll be in v0.55.0.
	a1.AltBech32M = penumbra_addr
	fmt.Println("Via protos, addr1 is:", a1)
	// Naively convert the string address into a byte slice. Abjure all knowledge of bech32m and similar encodings.
	a2.Inner = []byte(penumbra_addr)
	fmt.Println("Via protos, addr2 is:", a2)

	// Just as naively convert the byte slice address back into a string, creating a new Address struct.
	// We'll compare this to the original and ensure equality.
	a3.AltBech32M = string(a2.Inner[:])
	fmt.Println("Via protos, addr3 is:", a3)

	if reflect.DeepEqual(a1, a3) {
		fmt.Println("What good fortune: a1 & a3 are identical.")
	} else {
		fmt.Println("Tragedy has struck: a1 & a3 are not alike.")
	}

}

func main() {
	bech32m_round_trip()
	via_protos()
}

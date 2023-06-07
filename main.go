package main

import (
	"fmt"
	chia "github.com/chia-network/go-chia-libs/pkg/bech32m"
	"github.com/cosmos/btcutil/bech32"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pactus-project/pactus/util/bech32m"
)

func main() {

	penumbra_addr := "penumbrav2t1fc6gvmyz749cvf7qyz2cgqyw8aq2zf7cm005uvmt4l5dtew5xuvdtuvdrcpwn740xlgx9saeyqtqftwnw57q3vkyd73teckwm9jkwcmwcxml7q7klu9smekthxpa2575urjltu"
	p2 := ""
	fmt.Println("Penumbra address is:", penumbra_addr)
	hrp, as_b, err := bech32m.DecodeNoLimit(penumbra_addr)
	fmt.Println("Penumbra human-readable prefix:", hrp)
	fmt.Println("Penumbra address decoded to bytes looks like:", as_b)

	// Try via bech32m
	fmt.Println("Using bech32m nolimit conversion")
	p2, err = bech32m.EncodeFromBase256(hrp, as_b)
	if err != nil {
		fmt.Printf("Failed to convert addr from bytes via bech32m")
	}
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Println("As expected, the two addresses were identical round-trip.")
	}

	// Try via btcutil
	// This method is gets us the closest to least surprise: all but the final 6 bytes
	// of the second address are identical. That means it's the checksum that's different,
	// explicable by the fact we're using bech32 when we should be using bech32m.
	fmt.Println("Using btcutil conversion")
	p2, err = bech32.Encode(hrp, as_b)
	if err != nil {
		fmt.Printf("Failed to convert addr from bytes via btcutil")
	}
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Println("As expected, the two addresses were identical round-trip.")
	}

	// Try via cosmos-sdk
	fmt.Println("Using cosmos-sdk conversion")
	p2, err = sdk.Bech32ifyAddressBytes(hrp, as_b)
	if err != nil {
		fmt.Printf("Failed to convert addr from bytes via cosmos-sdk")
	}
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Println("As expected, the two addresses were identical round-trip.")
	}

	// Try via chia-bech32m
	fmt.Println("Using chia bech32m")
	hrp, as_b, _ = chia.Decode(penumbra_addr)
	p2 = chia.Encode(hrp, as_b)
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Println("As expected, the two addresses were identical round-trip.")
	}
}

package main

import (
	"context"
	"fmt"
	"io"
	"github.com/pactus-project/pactus/util/bech32m"
	// We import the local buf-built protos pulled from the penumbra repo so that
	// we have access to the new "AltBech32m" field on the Address proto struct.
	penumbracrypto_latest "penumbra/address/penumbraprotos/penumbra/core/crypto/v1alpha1"
	// penumbracrypto "github.com/strangelove-ventures/interchaintest/v7/chain/penumbra/core/crypto/v1alpha1"
	penumbracrypto "github.com/strangelove-ventures/interchaintest/v7/chain/penumbra/core/crypto/v1alpha1"
	penumbraview "github.com/strangelove-ventures/interchaintest/v7/chain/penumbra/view/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
)

// NOT WORKING. Intended to demonstrate how to convert a string representation of a Penumbra address
// into a base256-encoded byte slice of length 80, so that it can be passed via protos to pclientd.
// As of v0.55.0, we can just pass bech32m strings as a field in the Address proto struct, which
// will be a lot more ergonomic.
func bech32m_round_trip() {
	penumbra_addr := "penumbrav2t1fc6gvmyz749cvf7qyz2cgqyw8aq2zf7cm005uvmt4l5dtew5xuvdtuvdrcpwn740xlgx9saeyqtqftwnw57q3vkyd73teckwm9jkwcmwcxml7q7klu9smekthxpa2575urjltu"
	fmt.Println("Penumbra address is:", penumbra_addr)
	// hrp, type_byte, as_b, err := bech32m.DecodeToBase256WithTypeNoLimit(penumbra_addr)
	// hrp, as_b, err := bech32m.DecodeToBase256(penumbra_addr)
	// fmt.Println("Type byte came back as:", type_byte)
	hrp, as_b, err := bech32m.DecodeNoLimit(penumbra_addr)
	if err != nil {
		fmt.Println("Failed to decode string address to bytes via bech32m:", err)
	}
	fmt.Println("Penumbra human-readable prefix:", hrp)
	// This byte slice is suitable for passing back to `bech32m.Encode`, as done below,
	// but is *not* suitable for passing to gRPC methods like `AddressByIndexRequest`.
	fmt.Println("Penumbra address decoded to bytes looks like:", as_b)

	// Try via bech32m
	var p2 string
	fmt.Println("Using bech32m nolimit conversion")
	p2, err = bech32m.Encode(hrp, as_b)
	if err != nil {
		fmt.Println("Failed to convert encode bytes to string via bech32m")
	}
	if p2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	} else {
		fmt.Printf("As expected, the two addresses were identical round-trip::\n\t* %s\n\t* %s\n", penumbra_addr, p2)
	}
}

// Demonstrate how to convert an address byte slice, as received from pclientd,
// into a string representation of the same Penumbra address.
func via_pclientd() (bool, error) {
	// This penumbra address is different from the other test strings, because it must match the custody
	// file available locally. In interchaintest, this key material will be generated dynamically.
	// For this function to work, you must run pclientd locally at localhost:8081.
	penumbra_addr := "penumbrav2t1pc3h8704vrvpcrv4udx55qjlc62tw9epxz3s9sd0nu3w6w9d83euw0zrg388xexkqxcn8lyaf2rndu7ku8zk3yzsgcm02rs25nza60nsc47j29cq9y2kkha69aykd7h0ewr44u"
	pclientd_addr := "localhost:8081"
	fmt.Println("Dialing pclientd grpc address at", pclientd_addr)
	channel, err := grpc.Dial(
		pclientd_addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return false, err
	}
	defer channel.Close()

	viewClient := penumbraview.NewViewProtocolServiceClient(channel)
	addressReq := &penumbraview.AddressByIndexRequest{
		AddressIndex: &penumbracrypto.AddressIndex{
			Account: 0,
		}}

	ctx := context.TODO()
	fmt.Println("Requesting address via pclientd...")
	addressResponse, err := viewClient.AddressByIndex(ctx, addressReq)
	if err != nil {
		fmt.Println("Encountered error =(", err)
		return false, err
	}
	addrBytes := addressResponse.Address.Inner
	fmt.Println("We received a byte slice of length:", len(addrBytes))

	// Hardcode the human-readable part of the address
	hrp := "penumbrav2t"
	// Decode from Base256. Otherwise, the `bech32m.Encode` string will error out on values
	// in our slice greater than 32.
	penumbra_addr_2, err := bech32m.EncodeFromBase256(hrp, addrBytes)
	if err != nil {
		fmt.Println("Failed to encode address bytes as bech32m", err)
		return false, err
	}

	if penumbra_addr_2 != penumbra_addr {
		fmt.Printf("The two addresses are not the same:\n\t* %s\n\t* %s\n", penumbra_addr, penumbra_addr_2)
	} else {
		fmt.Printf("As expected, the two addresses were identical round-trip::\n\t* %s\n\t* %s\n", penumbra_addr, penumbra_addr_2)
	}

	// Use pclientd connection to look up balance info.
	balanceRequest := &penumbraview.BalanceByAddressRequest{
		Address: addressResponse.Address,
	}
	// The BalanceByAddress method returns a stream response, containing
	// zero-or-more balances.
	balanceStream, err := viewClient.BalanceByAddress(ctx, balanceRequest)
	var balances []penumbraview.BalanceByAddressResponse
	for {
		balance, err := balanceStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Failed to get balance: ", err)
				return false, err
			}
		}
		// fmt.Printf("Balance response looks like: %+q\n", balance)
		balances = append(balances, *balance)
	}
	fmt.Println("What do I have in my wallet? Behold:")
	for _, b := range balances {
		// N.B. the `Hi` and `Lo` fields on Amount denote high/low order bytes:
		// https://github.com/penumbra-zone/penumbra/blob/v0.54.1/crates/core/crypto/src/asset/amount.rs#L220-L240
		fmt.Printf("%v '%v'\n", b.Asset, b.Amount.Lo)
	}

	return true, nil
}

func via_protos() {
	penumbra_addr := "penumbrav2t1fc6gvmyz749cvf7qyz2cgqyw8aq2zf7cm005uvmt4l5dtew5xuvdtuvdrcpwn740xlgx9saeyqtqftwnw57q3vkyd73teckwm9jkwcmwcxml7q7klu9smekthxpa2575urjltu"
	var a1 penumbracrypto_latest.Address
	var a2 penumbracrypto_latest.Address
	var a3 penumbracrypto_latest.Address
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
	// via_protos()
	via_pclientd()
}

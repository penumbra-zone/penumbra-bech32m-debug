# penumbra-bech32m-debug

This repo is a scratchpad for wrangling Penumbra addresses in golang.
In short, naively decoding a Penumbra address string to bytes, then
re-encoding those bytes back to a stringified address, does not work reliably.
Let's figure out why.

## Running it
```
make
# or
go run .
```

## Takeaways

It's crucial to use bech32m, rather than bech32, format for decoding
and encoding.

## Relevant reading
* https://github.com/penumbra-zone/penumbra/blob/v0.53.1/crates/core/crypto/src/address.rs
* https://en.bitcoin.it/wiki/BIP_0350

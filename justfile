run:
    go run .

protos:
    buf generate --template proto/buf.gen.penumbra.yaml buf.build/penumbra-zone/penumbra

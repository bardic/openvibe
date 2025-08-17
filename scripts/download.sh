#! /bin/sh

cd ..

go run cmd/openvibe/openvibe.go download \
    --url "https://api.github.com/repos/Mojang/bedrock-samples/releases/latest"
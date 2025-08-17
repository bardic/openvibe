#!/bin/sh

cd ..

go run cmd/main/main.go config \
    --out "./openvibe-test.conf" \
    --build-name "open-vibe" \
    --name "openb vibe" \
    --description "vanilla pbr/vibrant texture pack" \
	--header-uuid "bd111387-103a-4ef4-b33b-b39af1270e50" \
    --module-uuid "5e867951-86e8-4c74-960c-1e7d1b8beedb" \
    --default-mer "[0,0,125,95]" \
	--version "[0,0,1]" \
    --author "judohippo" \
    --license "https://github.com/bardic/openpbr/blob/main/UNLICENSE" \
    --url "https://github.com/bardic/openpbr" \
	--capability "pbr" \
    --height-template "_height" \
    --mer-template "_mer" \
	--r-offset "1" \
    --g-offset "1" \
    --b-offset "1"
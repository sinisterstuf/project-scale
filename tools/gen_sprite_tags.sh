#!/bin/bash

if [[ $1 == "-h" ]]; then
	echo "this script extracts frame tags from an aseprite-exported sprite JSON and generates Go constants to refer to them by name" >&2
fi

if [[ $# -ne 3 ]]; then
	echo "this is script needs exactly 3 arguments: input file and output file and const prefix" >&2
	exit 1
fi

echo "Generating $3AnimationTags into $2 from sprite $1"

truncate -s 0 "$2"
echo -e "package main\n" >> "$2"
echo -e "// DO NOT EDIT\n// Generated by: $0\n" >> "$2"

echo -e "type $3AnimationTags uint8\n\nconst (" >> "$2"
jq -r '.meta.frameTags[].name' "$1" \
	| sed 's/[ \/]//g' \
	| sed "1s/$/ $3AnimationTags = iota/" \
	| sed "s/^/$3/" \
	>> "$2"
echo ")" >> "$2"

echo -e "var $3AnimationNames = []string {" >> "$2"
jq '.meta.frameTags[].name' "$1" | sed 's/$/,/' >> "$2"
echo "}" >> "$2"

gofmt -s -w "$2"

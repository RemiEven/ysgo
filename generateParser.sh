#!/bin/bash

docker pull ghcr.io/remieven/antlr4:latest

docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go internal/parser/YarnSpinnerLexer.g4
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go internal/parser/YarnSpinnerParser.g4

sed -i 's/GetText()/GetInnerText()/g' internal/parser/yarnspinner_parser.go
sed -i 's/\/\/ GetText/\/\/ GetInnerText/g' internal/parser/yarnspinner_parser.go

rm internal/parser/*.interp internal/parser/*.tokens

go fmt ./internal/parser/

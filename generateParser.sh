#!/bin/bash

docker pull ghcr.io/remieven/antlr4:latest

docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go parser/YarnSpinnerLexer.g4
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go parser/YarnSpinnerParser.g4

sed -i 's/GetText()/GetInnerText()/g' parser/yarnspinner_parser.go
sed -i 's/\/\/ GetText/\/\/ GetInnerText/g' parser/yarnspinner_parser.go

rm parser/*.interp parser/*.tokens

go fmt ./parser/

#!/bin/bash

docker build -f antlr/Dockerfile -t antlr/antlr4 .

docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work antlr/antlr4 -Dlanguage=Go parser/YarnSpinnerLexer.g4
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work antlr/antlr4 -Dlanguage=Go parser/YarnSpinnerParser.g4

sed -i 's/GetText()/GetInnerText()/g' parser/yarnspinner_parser.go
sed -i 's/\/\/ GetText/\/\/ GetInnerText/g' parser/yarnspinner_parser.go

rm parser/*.interp parser/*.tokens

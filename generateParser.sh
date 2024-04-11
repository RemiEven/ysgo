#!/bin/bash

# This script is used to generate the parser and listener of the internal/parser package with Antlr from the g4 files of YarnSpinner

set -e

docker pull ghcr.io/remieven/antlr4:latest

docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go internal/parser/YarnSpinnerLexer.g4
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work ghcr.io/remieven/antlr4 -Dlanguage=Go internal/parser/YarnSpinnerParser.g4

sed -i 's/(p \*YarnSpinnerLexer)/(l \*YarnSpinnerLexer)/g' internal/parser/yarnspinner_lexer.go
sed -i 's/IsInWhenClause()/l.inWhenClause/g' internal/parser/yarnspinner_lexer.go
sed -i 's/PushMode(/l.PushMode(YarnSpinnerLexer/g' internal/parser/yarnspinner_lexer.go
sed -i 's/LastTokenWas(ID/l.lastTokenWas(YarnSpinnerLexerID/g' internal/parser/yarnspinner_lexer.go
sed -i 's/SetInWhenClause/l.inWhenClause = /g' internal/parser/yarnspinner_lexer.go

sed -i 's/GetText()/GetInnerText()/g' internal/parser/yarnspinner_parser.go
sed -i 's/\/\/ GetText/\/\/ GetInnerText/g' internal/parser/yarnspinner_parser.go

rm internal/parser/*.interp internal/parser/*.tokens

go fmt ./internal/parser/

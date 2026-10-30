#!/bin/bash
alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:.:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -package parser -visitor *.g4

#windows
# java -Xmx500M -cp "antlr-4.13.0-complete.jar;." org.antlr.v4.Tool -Dlanguage=Go -package parser -visitor *.g4
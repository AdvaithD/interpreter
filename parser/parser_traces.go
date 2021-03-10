package parser

import (
        "fmt" 
        "strings"
)

var traceLevel int = 0
const traceIdentPlaceholder string = "\t"

func identLevel() string {
    return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(string string) {
    fmt.Printf("%s%s\n", identLevel(), string)
}

func incIdent() { traceLevel = traceLevel + 1}
func decIdent() { traceLevel = traceLevel - 1}

func trace(msg string) string {
    incIdent()
    tracePrint("BEGIN " + msg)
    return msg
}

func untrace(msg string) {
    tracePrint("ENG " + msg)
    decIdent()
}
 

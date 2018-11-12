package models

import (
    "testing"
    "fmt"
)


func TestParserSelection(t *testing.T) {

    parser := NewEngineParser("../engines/Estes_A8.rse")
    equals(t, "../engines/Estes_A8.rse", parser.FileName)
    parser = parser.Parse("")
    fmt.Printf("Datapoints are %v\n", parser.Data)
}
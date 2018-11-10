package models_tests

import (
    "testing"
    "github.com/mwhagedorn/gorocksim/models"
    "fmt"
)


func TestParserSelection(t *testing.T) {

    parser := models.NewEngineParser("../engines/Estes_A8.rse")
    equals(t, "../engines/Estes_A8.rse", parser.FileName)
    parser.Parse()
    fmt.Printf("Datapoints are %v\n", parser.Data)
}
package models

import (
	"path/filepath"
)


type EngineParser struct {
	FileName string
	Processor Parser 
	Code string
	Diameter float64
	Length float64
	Delay string
	PropWeight float64
	EngineWeight float64
	Manufacturer string
	BurnTime float64
	Data []RSEDataPoint
}

func NewEngineParser(filename string) EngineParser {
    var p EngineParser
    p.FileName = filename
	p.Processor = new_parser(filename)
    return p
}

func new_parser(filename string) Parser {
	if filepath.Ext(filename) == ".eng" {
		var p WraspParser
		return &p
	}
	
	if filepath.Ext(filename) == ".rse" {
		var p RSEParser
		return &p
	}
	
	return nil
}

func (e EngineParser) Parse(filename string) Parser {
	return e.Processor.Parse(e.FileName)
}


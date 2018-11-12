package models

type Parser interface {
	Parse(filename string) EngineData
}

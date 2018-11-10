package models

type Parser interface {
	parse(filename string) *Parser
}
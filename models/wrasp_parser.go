package models
import (
	"fmt"
)

type WraspParser struct {
	Context EngineParser
}

func (w WraspParser) Parse(filename string) Parser{
	fmt.Printf("parse()\n")
	return w.Context
}

func (w WraspParser) setContext(parent EngineParser)  {
	w.Context = parent
}

func (w WraspParser) force_value_at(time float64) float64 {
	return 0.0
}
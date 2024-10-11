package facts

import (
	"fmt"
	"time"
)

// MyFact is now exported (capitalized)
type MyFact struct {
	IntAttribute     int64
	StringAttribute  string
	BooleanAttribute bool
	FloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

// GetWhatToSay is now exported (capitalized)
func (mf *MyFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}

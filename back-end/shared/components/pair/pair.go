package Pair

import "strings"

type Pair struct {
	Base  string
	Quote string
}

func (p *Pair) Symbol() string {
	return p.Base + "-" + p.Quote
}

func (p *Pair) StringLowercase() string {
	return strings.ToLower(p.Base) + "-" + strings.ToLower(p.Quote)
}

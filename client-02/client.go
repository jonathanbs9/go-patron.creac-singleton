package client_02

import "github.com/jonathanbs9/go-patron.creac-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}

package main

import (
	"fmt"
	"sync"

	client_01 "github.com/jonathanbs9/go-patron.creac-singleton/client-01"
	client_02 "github.com/jonathanbs9/go-patron.creac-singleton/client-02"
	"github.com/jonathanbs9/go-patron.creac-singleton/singleton"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_01.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			client_02.IncrementAge()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}

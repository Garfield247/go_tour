package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (this *Name) String() string {
	return fmt.Sprint(*this)
}

func (this *Name) Set(value string) error {
	if len(*this) > 0 {
		return errors.New("name flag already set")

	}
	*this = Name("catman: " + value)
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)

}

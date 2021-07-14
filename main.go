package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	log.Printf("args: %v", args)
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "GO 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "python":
		pyCmd := flag.NewFlagSet("python", flag.ExitOnError)
		pyCmd.StringVar(&name, "n", "python 语言", "帮助信息")
		_ = pyCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}

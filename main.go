package main

import (
	"flag"
	"fmt"
	"os"
	"plate/cmd"
)

func main() {
	template := flag.String("template", "", "Choose a template structure CMD, API")
	projectName := flag.String("name", "", "name of the project")
	flag.Parse()
	if *template == "" {
		fmt.Println("Choose a template structure CMD, API")
		flag.Usage()
		os.Exit(1)
	}
	if *projectName == "" {
		fmt.Println("Set the name of the project")
		flag.Usage()
		os.Exit(1)
	}

	if err := cmd.Execute(*template, *projectName); err != nil {
		if err == flag.ErrHelp {
			os.Exit(2)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

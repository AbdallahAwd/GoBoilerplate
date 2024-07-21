package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AbdallahAwd/GoBoilerplate/cmd"
	"github.com/AbdallahAwd/GoBoilerplate/pkg/shared"
)

func main() {
	template := flag.String("template", "", "Choose a template structure cli, api, all")
	projectName := flag.String("name", "", "name of the project")
	showVersion := flag.Bool("version", false, "Show the version of the tool")

	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", shared.AppVersion)
		os.Exit(0)
	}
	if *template == "" {
		fmt.Println("Choose a template structure cli, api, General Web App")
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

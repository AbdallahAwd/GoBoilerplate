package cmd

import (
	"flag"
	"fmt"

	"github.com/AbdallahAwd/plate/internal/commands"
)

func Execute(temp string, projectName string) error {

	switch temp {
	case "cli":
		return commands.Cmd(projectName)
	case "all":
		return commands.UltimateProjectStructure(projectName)
	case "api":
		return commands.APIProject(projectName)

	default:
		flag.Usage()
		return fmt.Errorf("unknown command")
	}
}

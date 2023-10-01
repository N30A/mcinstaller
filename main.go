package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, Usage)
		os.Exit(0)
	}

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)

	if len(os.Args) < 2 {
		flag.Usage()
	}

	switch os.Args[1] {

	case "list":
		listCmd.Parse(os.Args[2:])
		ListSubcommand(listCmd.Args())

	case "install":
		installCmd.Parse(os.Args[2:])
		InstallSubcommand(installCmd.Args())

	case "help":
		flag.Usage()

	default:
		fmt.Println("Invalid command - use 'help' to see the help menu")
		os.Exit(1)
	}
}

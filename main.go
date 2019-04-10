package main

import (
	"github.com/laginha87/cstore/cstore/actions"
	"github.com/spf13/cobra"
)

func main() {

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Inits it",
		Run:   actions.Init,
	}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add command",
		Run:   actions.Add,
	}
	var resetCmd = &cobra.Command{
		Use:     "reset",
		Short:   "Reset",
		Run:     actions.Reset,
		Aliases: []string{"r"},
	}

	var execCmd = &cobra.Command{
		Use:     "exec",
		Run:     actions.Exec,
		Aliases: []string{"e"},
	}

	var listCmd = &cobra.Command{
		Use:     "list",
		Run:     actions.List,
		Aliases: []string{"l"},
	}

	var zshCmd = &cobra.Command{
		Use: "zsh",
		Run: actions.Zsh,
	}
	var rootCmd = &cobra.Command{Use: "command-storer"}
	rootCmd.AddCommand(initCmd, addCmd, resetCmd, listCmd, execCmd, zshCmd)
	rootCmd.Execute()
}

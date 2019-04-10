package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Zsh Shows alias to add to a zsh init file
func Zsh(command *cobra.Command, args []string) {
	fmt.Println("Add the following to your zsh file:")
	fmt.Println("")
	fmt.Println("alias c1=cstore exec 1")
	fmt.Println("alias c2=cstore exec 2")
	fmt.Println("alias c3=cstore exec 3")
	fmt.Println("alias c4=cstore exec 4")
	fmt.Println("alias c5=cstore exec 5")
	fmt.Println("alias cl=cstore list")
	fmt.Println("alias ca=cstore add")
}

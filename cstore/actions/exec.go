package actions

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/laginha87/cstore/cstore"
	"github.com/spf13/cobra"
)

// Exec exec action
func Exec(c *cobra.Command, args []string) {
	lineno, _ := strconv.Atoi(args[0])
	cmd := fmt.Sprintf("%s", cstore.GetNthLine(cstore.FilePath, lineno))
	syscall.Exec("/bin/zsh", []string{"zsh", "-c", "-l", cmd}, os.Environ())
}

package actions

import (
	"os"

	"github.com/laginha87/cstore/cstore"
	"github.com/spf13/cobra"
)

// Reset reset action
func Reset(cmd *cobra.Command, args []string) {
	os.OpenFile(cstore.FilePath, os.O_TRUNC, 0666)
}

package actions

import (
	"os"

	"github.com/laginha87/cstore/cstore"
	"github.com/spf13/cobra"
)

// Init init action
func Init(cmd *cobra.Command, args []string) {
	os.OpenFile(cstore.FilePath, os.O_RDONLY|os.O_CREATE, 0666)
}

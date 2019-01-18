package actions

import (
	"os"
	"strings"

	"github.com/laginha87/cstore/cstore"
	"github.com/spf13/cobra"
)

// Add add action
func Add(command *cobra.Command, args []string) {
	var lastLine = cstore.GetLastLineWithSeek(cstore.HistPath)
	var options = strings.Split(lastLine, ";")
	var cmd = options[len(options)-1]

	f, err := os.OpenFile(cstore.FilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(cmd); err != nil {
		panic(err)
	}
}

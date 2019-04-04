package actions

import (
	"bufio"
	"fmt"
	"os"

	"github.com/laginha87/cstore/cstore"

	"github.com/spf13/cobra"
)

// List list action
func List(cmd *cobra.Command, args []string) {
	fileHandle, err := os.Open(cstore.FilePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	sc := bufio.NewScanner(fileHandle)
	i := 1
	for sc.Scan() {
		fmt.Printf("%d: %s\n", i, sc.Text())
		i++
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var damperCmd = &cobra.Command{
		Use:   "damper",
		Short: "Damper Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	err := damperCmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/ayuxsec-org/log"
)

func main() {
	fmt.Fprint(os.Stderr, banner)
	rootCmd := NewCmdi().RootCmd()
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("'rootCmd.Execute' error: %v", err)
	}
}

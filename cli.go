package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/odlp/inflight/runner"
)

const currentVersion = "0.2.0"

func main() {
	var displayVersion = flag.Bool("version", false, "view version information")
	var outputPath = flag.String("o", "", "output location for commit message")
	flag.Parse()

	if displayVersion != nil && *displayVersion {
		fmt.Println(currentVersion)
		os.Exit(0)
	}

	if outputPath == nil || *outputPath == "" {
		fmt.Println("Inflight: Output location for commit message required")
		os.Exit(1)
	}

	r := runner.NewRunner(*outputPath)
	r.Exec()
}

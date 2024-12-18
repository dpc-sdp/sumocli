package main

import (
	"github.com/dpc-sdp/sumocli/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}

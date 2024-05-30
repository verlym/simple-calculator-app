/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycalc",
	Short: "mycalc - a commandline calculator for basic arithmetic + - * /",
	Long: `mycalc is a simple yet powerful command line personal calculator

	You can use mycalc for quick calculations at the command line.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/verlym/mycalculator-cli/pkg/mycalc"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"m"},
	Short:   "Multiplies a bunch of integers",
	Run: func(cmd *cobra.Command, args []string) {
		res := mycalc.Multiply(args)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiplyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

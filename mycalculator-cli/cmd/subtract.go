/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/verlym/mycalculator-cli/pkg/mycalc"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"s"},
	Short:   "Subtracts two integers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := mycalc.Subtract(args)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

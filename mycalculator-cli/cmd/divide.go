/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/verlym/mycalculator-cli/pkg/mycalc"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"d"},
	Short:   "Divides two integers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := mycalc.Divide(args)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

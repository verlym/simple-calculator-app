/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/verlym/mycalculator-cli/pkg/mycalc"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Adds a bunch of integers",
	Run: func(cmd *cobra.Command, args []string) {
		res := mycalc.Add(args)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

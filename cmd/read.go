/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read command - reads from config file",
	Long:  `Read from config file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n\n", viper.AllKeys())

		settings := viper.AllSettings()

		for i, value := range settings {
			fmt.Printf("%v: %v\n", i, value)
		}

		// viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
		foo, _ := cmd.PersistentFlags().GetString("foo")
		fmt.Println(foo)

		toggle, _ := cmd.Flags().GetBool("toggle")
		fmt.Println(toggle)
	},
}

func init() {
	fmt.Printf("running from read.go\n\n")
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

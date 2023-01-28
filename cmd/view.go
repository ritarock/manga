/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ritarock/manga/internal/viewer"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view manga",
	Long: `This subcommand is used to view data.
This command execute it will start the server. (http://localhost:8080/manga)

You can specify a query path. Default is the date of execution.
ex) http://localhost:8080/manga?yyyy=2022&mm=1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start server: http://localhost:8080/manga")
		viewer.Run()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

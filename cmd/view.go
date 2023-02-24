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
This command execute it will start the server.

You can specify a query path. Default is the month of execution.
ex) http://localhost:8080/manga?yyyy=2022&mm=1

You can also use GraphQL playground.
http://localhost:8080/playground
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", `Start Server: http://localhost:8080/manga
GraphQL playground: http://localhost:8080/playground
`)
		viewer.Run()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

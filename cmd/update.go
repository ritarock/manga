/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ritarock/manga/internal/crawler"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update manga data",
	Long: `This subcommand is used to update manga data.
This subcommand must be executed first.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := crawler.Run(); err != nil {
			return nil
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

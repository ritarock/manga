package cmd

import (
	"context"

	"github.com/ritarock/manga/internal/crawler"
	"github.com/ritarock/manga/internal/db"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update manga data",
	Long: `This subcommand is used to update manga data.
This subcommand must be executed first.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.InitDb()
		client, err := db.Connection()
		if err != nil {
			return err
		}
		defer client.Close()

		ctx := context.Background()
		client.Book.Delete().Exec(ctx)

		coverages := crawler.GetCoverages()
		isbnList := crawler.MakeIsbnList(coverages)
		for _, isbn := range isbnList {
			books := crawler.GetBooks(isbn)
			for _, book := range books {
				if err := db.Store(ctx, client, book); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

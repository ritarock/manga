package cmd

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/ritarock/manga/internal/db"
	"github.com/ritarock/manga/internal/viewer"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view manga",
	Long: `This subcommand is used to view data.
This command execute it will start the server.

You can specify a query path. Default is the month of execution.
ex) http://localhost:8080/manga?yyyy=2022&mm=1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start server: http://localhost:8080/manga")

		server := http.Server{
			Addr: "0.0.0.0:8080",
		}
		http.HandleFunc("/manga", index)
		server.ListenAndServe()
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	var yyyy, mm string
	if r.URL.Query().Has("yyyy") {
		yyyy = viewer.ValidateYyyy(r.FormValue("yyyy"))
	}
	if r.URL.Query().Has("mm") {
		mm = viewer.ValidateMm(r.FormValue("mm"))
	}
	if yyyy == "" || mm == "" {
		yyyy, mm = func(t time.Time) (string, string) {
			return t.Format("2006"), t.Format("01")
		}(time.Now())
	}
	ctx := context.Background()
	client, err := db.Connection()
	if err != nil {
		panic(err)
	}
	books, err := db.GetByDate(ctx, client, yyyy, mm)
	if err != nil {
		panic(err)
	}
	var viewTemplate []viewer.ViewTemplate
	for _, book := range books {
		viewTemplate = append(viewTemplate, viewer.ViewTemplate{
			Cover: book.Cover,
			Title: book.Title,
		})
	}
	tmpl, err := template.ParseFiles("view/tmpl.html")
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, viewTemplate); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

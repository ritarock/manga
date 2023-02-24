package viewer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ritarock/manga/ent"
	"github.com/ritarock/manga/ent/book"
	"github.com/ritarock/manga/graph"
	"github.com/ritarock/manga/internal/db"
)

type ViewTemplate struct {
	Cover string
	Title string
}

const viewTemplatePath = "view/tmpl.html"

func Run() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	client, _ := db.Connection()
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{
				EntClient: client,
			},
			},
		),
	)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/manga", index)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	var yyyy, mm string
	if r.URL.Query().Has("yyyy") {
		yyyy = validateYyyy(r.FormValue("yyyy"))
	}
	if r.URL.Query().Has("mm") {
		mm = validateMm(r.FormValue("mm"))
	}
	if yyyy == "" || mm == "" {
		yyyy, mm = func(t time.Time) (string, string) {
			return t.Format("2006"), t.Format("01")
		}(time.Now())
	}
	bookCovers := getBooks(yyyy, mm)
	tmpl, err := template.ParseFiles(viewTemplatePath)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, bookCovers); err != nil {
		panic(err)
	}
}

func getBooks(yyyy, mm string) []ViewTemplate {
	client, _ := db.Connection()
	defer client.Close()

	// SELECT Cover, Title FROM books WHERE (Pubdate LIKE "YYYYMM%" AND Cover != "") ORDER BY Pubdate DESC;
	books, err := client.Book.Query().
		Select(book.FieldCover, book.FieldTitle).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.Like(book.FieldPubdate, yyyy+mm+"%"))
			},
			book.Not(book.CoverEQ("")),
		).
		Order(ent.Desc(book.FieldPubdate)).
		All(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	var viewTemplate []ViewTemplate
	for _, book := range books {
		viewTemplate = append(viewTemplate, ViewTemplate{
			Cover: book.Cover,
			Title: book.Title,
		})
	}

	return viewTemplate
}

func validateYyyy(yyyy string) string {
	i, err := strconv.Atoi(yyyy)
	if err != nil {
		return ""
	}
	if i >= 1970 {
		return yyyy
	}
	return ""
}

func validateMm(mm string) string {
	i, err := strconv.Atoi(mm)
	if err != nil {
		return ""
	}
	if 1 <= i && i <= 12 {
		return fmt.Sprintf("%02d", int64(i))
	}

	return ""
}

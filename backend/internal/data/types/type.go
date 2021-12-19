package types

type Book struct {
	Isbn        string `db:"isbn"`
	Title       string `db:"title"`
	Publisher   string `db:"publisher"`
	Pubdate     string `db:"pubdate"`
	Cover       string `db:"cover"`
	Author      string `db:"author"`
	SubjectCode string `db:"subject_code"`
}

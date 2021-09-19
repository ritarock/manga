package data

type Book struct {
	Isbn        string `db:"isbn" json:"isbn"`
	Title       string `db:"title" json:"title"`
	Publisher   string `db:"publisher" json:"publisher"`
	Pubdate     string `db:"pubdate" json:"pubdate"`
	Cover       string `db:"cover" json:"cover"`
	Author      string `db:"author" json:"author"`
	SubjectCode string `db:"subject_code" json:"SubjectCode"`
}

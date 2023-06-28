package types

type OpenBD []struct {
	Onix struct {
		DescriptiveDetail struct {
			Subject []struct {
				SubjectCode string `json:"SubjectCode"`
			}
		} `json:"DescriptiveDetail"`
	} `json:"onix"`
	Summary struct {
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Publisher string `json:"publisher"`
		Pubdate   string `json:"pubdate"`
		Cover     string `json:"cover"`
		Author    string `json:"author"`
	} `json:"summary"`
}

type Book struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Pubdate     string `json:"pubdate"`
	Cover       string `json:"cover"`
	Author      string `json:"author"`
	SubjectCode string `json:"SubjectCode"`
}

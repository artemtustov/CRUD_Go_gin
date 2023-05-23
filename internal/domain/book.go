package domain

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
	Count     int    `json:"count"`
}

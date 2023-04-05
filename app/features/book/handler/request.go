package handler

type BookRequest struct {
	Judul string `json:"judul"`
	Tahun string `json:"tahun"`
	Penerbit string `json:"penerbit"`
}
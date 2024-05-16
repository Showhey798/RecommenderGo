package entity

type Movie struct {
	ID     uint32 `json:"id"`
	Title  string `json:"title"`
	ImgUrl string `json:"img_url"`
}

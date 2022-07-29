package book

type Input struct {
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
	//SubTitle string      `json:"sub_title"`
}
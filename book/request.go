package book

type Request struct {
	Title       string      `json:"title" binding:"required"`
	Price       interface{} `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"required,number"`
	Discount    int         `json:"discount" binding:"required,number"'`
	//SubTitle string      `json:"sub_title"`
}

package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest Request) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest Request) (Book, error) {
	var price int
	_, status := bookRequest.Price.(float64)
	if status {
		price = int(bookRequest.Price.(float64))
	} else {
		price = bookRequest.Price.(int)
	}

	book := Book{
		Title:       bookRequest.Title,
		Price:       price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

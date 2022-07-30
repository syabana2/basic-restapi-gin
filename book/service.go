package book

import "basic-rest-api-gin/helper"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest Request) (Book, error)
	Update(ID int, bookRequest Request) (Book, error)
	Delete(ID int) (Book, error)
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
	price := helper.ConvertInterfaceToInt(bookRequest.Price)
	rating := helper.ConvertInterfaceToInt(bookRequest.Rating)
	discount := helper.ConvertInterfaceToInt(bookRequest.Discount)

	book := Book{
		Title:       bookRequest.Title,
		Price:       price,
		Description: bookRequest.Description,
		Rating:      rating,
		Discount:    discount,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest Request) (Book, error) {
	book, err := s.repository.FindByID(ID)
	helper.FatalIfError(err)

	price := helper.ConvertInterfaceToInt(bookRequest.Price)
	rating := helper.ConvertInterfaceToInt(bookRequest.Rating)
	discount := helper.ConvertInterfaceToInt(bookRequest.Discount)

	book.Title = bookRequest.Title
	book.Price = price
	book.Description = bookRequest.Description
	book.Rating = rating
	book.Discount = discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	helper.FatalIfError(err)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}

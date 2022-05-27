package repository

import (
	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
)

type RepositoryInteractorRespository interface {
	SaveBook(book inventory.Book) error
	//GetBook(isbn string) (inventory.Book, error)
	ListInventory() ([]inventory.Book, error)
	//DeleteBook(isbn string) error
}

type RepositoryInteractor struct {
	handler RepositoryInteractorRespository
}

func NewRepositoryInteractor(handler RepositoryInteractorRespository) *RepositoryInteractor {
	return &RepositoryInteractor{handler}
}

func (r RepositoryInteractor) SaveBook(book inventory.Book) error {
	return r.handler.SaveBook(book)
}

// func (r RepositoryInteractor) GetBook(isbn string) (inventory.Book, error) {
// 	return r.handler.GetBook(isbn)
// }

func (r RepositoryInteractor) ListInventory() ([]inventory.Book, error) {
	results, _ := r.handler.ListInventory()
	return results, nil
}

// func (r RepositoryInteractor) DeleteBook(isbn string) error {
// 	return r.handler.DeleteBook(isbn)
// }

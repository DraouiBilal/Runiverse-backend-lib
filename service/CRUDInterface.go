package service

type CRUDInterface[T any] interface {
	GetAll() ([]T, error)
	GetById(id string) (T, error)
	Create(item T) (T, error)
	Update(item T) (T, error)
	Delete(id string) error
}

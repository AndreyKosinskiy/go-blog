package repository

type Repository interface {
	Create()
	Update()
	Delete()
	getById()
}

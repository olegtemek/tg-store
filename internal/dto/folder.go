package dto

type FolderCreate struct {
	Title  string `validate:"required"`
	UserId int    `validate:"required"`
}

type FolderUpdate struct {
	Id     int    `validate:"required"`
	Title  string `validate:"required"`
	UserId int    `validate:"required"`
}

type FolderGetAll struct {
	UserId int `validate:"required"`
}

type FolderGetOne struct {
	Id     int `validate:"required"`
	UserId int `validate:"required"`
}

type FolderDelete struct {
	Id     int `validate:"required"`
	UserId int `validate:"required"`
}

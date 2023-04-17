package schemas

import (
	"github.com/Simo-C3/tyoukankakuC3-2023/api/models"
)

type ToDos struct {
	ToDos []models.ToDo `json:"todos"`
}

type ToDo struct {
	ToDos models.ToDo `json:"todo"`
}

type PostToDo struct {
	Title  string            `gorm:"not null" json:"title"`
	Detail string            `gorm:"not null" json:"detail"`
	Status models.BaseStatus `gorm:"not null" json:"status"`
}

type PatchToDo struct {
	Title  *string            `json:"title"`
	Detail *string            `json:"detail"`
	Status *models.BaseStatus `json:"status"`
}

package store

import (
	"strconv"

	"github.com/aut-ce/students.go/model"
	"gorm.io/gorm"
)

type Student interface {
	Save(model.Student) error
	Load(int) (model.Student, error)
}

type SQLStudent struct {
	db *gorm.DB
}

func NewSQLStudent(db *gorm.DB) Student {
	return &SQLStudent{
		db: db,
	}
}

func (sql *SQLStudent) Save(s model.Student) error {
	result := sql.db.Create(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sql *SQLStudent) Load(id int) (model.Student, error) {
	var s model.Student

	result := sql.db.Where("id = ?", strconv.Itoa(id)).First(&s)
	if result.Error != nil {
		return model.Student{}, result.Error
	}

	return s, nil
}

package models

import (
	"github.com/nid90/kawaii-blog-engine/database"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
	Nick      string
	Avatar    string
	Email     string
	Password  string
}

func FindAuthorByEmail(email string) (*Author, error) {
	var author Author

	if err := database.DBConn.Where(&Author{Email: email}).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &author, nil
}
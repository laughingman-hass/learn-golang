package models

import "github.com/jinzhu/gorm"

func first(db *gorm.DB, record interface{}) error {
	err := db.First(record).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}

package service

import (
	"../model"
	"gorm.io/gorm"
)

func TaskPersist(owner string, task *model.Task) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var user model.User

		if err := tx.Where("username = ?", owner).First(&user).Error; err != nil {
			return err
		}

		task.UserID = user.ID

		if err := tx.Save(task).Error; err != nil {
			return err
		}

		return nil
	})
}

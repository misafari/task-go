package service

import "../model"

func UserInsert(user *model.User)  {
	DB.Save(user)
}

func UserDelete(user *model.User) {
	DB.Delete(user)
}

func UsersFindAll(users *[]model.User) {
	DB.Find(users)
}

func UserFindByUsername(username string, user *model.User) error {
	return DB.Where("username = ?", username).First(user).Error
}
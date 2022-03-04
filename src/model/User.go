package Models

import(
	"fmt"
	"gin-rest-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]User) (err error){
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//user memory location was passed into function
// can only return error or nil (where nil means ops was successfult)
//how do we get user data? by the referenced model

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error){
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
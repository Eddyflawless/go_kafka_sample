package Controllers

import(
	"fmt"
	"net/http"
	"gin-rest-api/Models"
	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context){

	var user []Models.User

	err := Models.GetAllUsers(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}

}


func CreateUser(c *gin.Context){

	var user Models.User

	c.BindJSON(&user)

	err := Models.CreateUser(&user)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
	
}

func GetUserByID(c *gin.Context){
	
}

func UpdateUser(c *gin.Context){

}

func DeleteUser(c *gin.Context){

}
package Routes

import (
	"gin-rest-api/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/v1")
	{
		grp1.GET("users", Controllers.GetUsers)
		grp1.POST("usesr", Controllers.CreateUser)
		grp1.GET("users/:id", Controllers.GetUserByID)
		grp1.PUT("users/:id", Controllers.UpdateUser)
		grp1.DELETE("users/:id", Controllers.DeleteUser)
	} 

	return r
}
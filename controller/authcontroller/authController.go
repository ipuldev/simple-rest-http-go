package authcontroller

import(
	_ "github.com/appleboy/gin-jwt/v2"
	"github.com/sobatfillah/config"
	"github.com/sobatfillah/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context){
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	var userCheck model.User
	config.GetDB().First(&userCheck,"username = ?", user.Username)

	if userCheck.ID > 0{
		c.JSON(http.StatusConflict, gin.H{"error" : "User Already Exists"})
		return
	}

	config.GetDB().Save(&user)

	c.JSON(http.StatusCreated,gin.H{"message" : "User Created Successfully"})
}
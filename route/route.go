package route


import(
	"github.com/sobatfillah/controller/authcontroller"
	"github.com/sobatfillah/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Routes() *gin.Engine{
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	if err != nil{
		log.Fatal("JWT Error:" + err.Error())
	}

	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Welcome to Service Sobatfillah")
	})

	v1 := router.Group("/v1")

	{
		v1.POST("/register",authcontroller.Register)
	}

	authorization := router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)

	return router
}
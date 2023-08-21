package restdelivery

import (
	"net/http"

	"github.com/aszanky/gofolderingproject/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Delivery struct {
	usecase usecase.Usecase
	router  *gin.Engine
}

func NewDelivery(
	uc usecase.Usecase,
) *Delivery {
	r := gin.Default()
	return &Delivery{
		usecase: uc,
		router:  r,
	}
}

func (d *Delivery) Register() {
	d.router.GET("/", d.PING)

	//Token
	// d.router.POST("/token", d.GenerateToken)

	//Users
	// d.router.POST("/create_user", d.CreateUser)
	// d.router.GET("/get_user", d.GetUser)
}

// Start HTTP Server
func (d *Delivery) Start(address string) error {
	return d.router.Run(address)
}

func (d *Delivery) PING(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "PONG",
	})
}

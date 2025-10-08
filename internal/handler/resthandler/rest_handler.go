package resthandler

import (
	"net/http"

	"github.com/aszanky/gofolderingproject/config"
	"github.com/aszanky/gofolderingproject/internal/usecase"
	"github.com/gin-gonic/gin"
)

type handler struct {
	usecase usecase.Usecase
	cfg     *config.MainConfig
}

func NewHandler(
	usecase usecase.Usecase,
	cfg *config.MainConfig,
) *handler {
	return &handler{
		usecase: usecase,
		cfg:     cfg,
	}
}

// Start HTTP Server
func (d *handler) Start(address string) error {
	return d.router.Run(address)
}

func (d *handler) PING(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "PONG",
	})
}

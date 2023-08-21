package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aszanky/gofolderingproject/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// Auth Extract Token from Authorization Header
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				c.JSON(http.StatusBadRequest, gin.H{
					"error":       err,
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")
		if len(idTokenHeader) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("please provide token"),
			})
			c.Abort()
			return
		}

		// validate ID token here
		user, err := auth.ValidateToken(idTokenHeader[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New("token invalid"),
			})
			c.Abort()
			return
		}

		c.Set("username", user)

		c.Next()
	}
}

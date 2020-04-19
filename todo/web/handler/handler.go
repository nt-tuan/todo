package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//Invalidate model
type Invalidate struct {
	Key   string `json:"field"`
	Value string `json:"message"`
}

//GetError invalidated model
func GetError(err error) *[]Invalidate {
	rs := make([]Invalidate, 0, 0)
	if err == nil {
		return &rs
	}
	switch err := err.(type) {
	case validator.ValidationErrors:
		for _, err := range err {
			rs = append(rs, Invalidate{err.Field(), err.Tag()})
		}
	default:
		rs = append(rs, Invalidate{"_", err.Error()})
	}
	return &rs
}

//JSONError default json error
func JSONError(c *gin.Context, err error) error {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, GetError(err))
	}
	return err
}

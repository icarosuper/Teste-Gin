package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ParseRequestBody[T any](context *gin.Context) (T, error) {
	var p T

	err := json.NewDecoder(context.Request.Body).Decode(&p)

	if err != nil {
		return p, err
	}

	return p, nil
}

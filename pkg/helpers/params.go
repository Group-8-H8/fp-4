package helpers

import (
	"strconv"

	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context, key string) (int, errs.Errs) {
	val := c.Param(key)

	id, err := strconv.Atoi(val)

	if err != nil {
		return 0, errs.NewBadRequestError("invalid parameter id")
	}

	return id, nil
}
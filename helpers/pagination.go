package helpers

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type paginationQuery struct {
	page  int
	limit int
}

func Pagination(c *gin.Context) (*options.FindOptions, error) {
	findOptions := options.Find()

	page, _ := strconv.ParseInt(c.Request.URL.Query().Get("page"), 10, 32)
	limit, _ := strconv.ParseInt(c.Request.URL.Query().Get("limit"), 10, 32)

	if reflect.TypeOf(page).String() != "int64" || reflect.TypeOf(limit).String() != "int64" {
		return nil, errors.New("Invalid parameters values!")
	}

	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	return findOptions, nil
}

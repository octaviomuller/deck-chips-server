package helpers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type paginationQuery struct {
	page  int
	limit int
}

func Pagination(c *gin.Context, FindOptions *options.FindOptions) {
	page, _ := strconv.ParseInt(c.Params.ByName("page"), 10, 32)
	limit, _ := strconv.ParseInt(c.Params.ByName("limit"), 10, 32)

	if reflect.TypeOf(page).String() != "int32" || reflect.TypeOf(limit).String() != "int32" {
		c.JSON(http.StatusUnprocessableEntity, ResponseMessage("Invalid parameters values!"))
	}

	FindOptions.SetSkip((page - 1) * limit)
	FindOptions.SetLimit(limit)
}

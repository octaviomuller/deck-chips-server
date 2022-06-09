package helper

import (
	"errors"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(page string, limit string) (*options.FindOptions, error) {
	findOptions := options.Find()

	pageValue, _ := strconv.ParseInt(page, 10, 32)
	limitValue, _ := strconv.ParseInt(limit, 10, 32)

	if reflect.TypeOf(pageValue).String() != "int64" || reflect.TypeOf(limitValue).String() != "int64" {
		return nil, errors.New("Invalid parameters values!")
	}

	findOptions.SetSkip((pageValue - 1) * limitValue)
	findOptions.SetLimit(limitValue)

	return findOptions, nil
}

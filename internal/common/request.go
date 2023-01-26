package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	fieldID = "id"
)

func BindIds(g gin.Context) (result []int) {
	//parse id param
	if id, err := strconv.Atoi(g.Param(fieldID)); err == nil {
		result = append(result, id)
	}

	// parse list of param
	params := g.Request.URL.Query()
	ids, ok := params[fieldID]

	if !ok {
		return result
	}

	for _, id := range ids {
		if id, err := strconv.Atoi(id); err == nil {
			result = append(result, id)
		}
	}

	return result
}

package common

import (
	"encoding/json"
	"errors"
	pagination_const "parishioner_management/internal/constant/pagination"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	fieldID = "id"
)

const (
	paramPage       = "page"
	paramFilter     = "filters"
	paramLimit      = "limit"
	paramSort       = "sorts"
	paramAscending  = "ASC"
	paramDescending = "DESC"
)

type SortParam struct {
	Field     string `json:"field"`
	Ascending bool   `json:"ascending"`
}

type ConditionalParam struct {
	Op     string   `json:"op"`
	Field  string   `json:"field"`
	Values []string `json:"val"`
}

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
}

type Param struct {
	// Page defines how should the server do pagination
	Page Paging `json:"page"`

	// Filters are the list of ConditionalParam that server should use filter data
	Filters []ConditionalParam `json:"filters"`

	// Sorts are list of SortParam in their priority orders that server will use
	// to sort result.
	Sorts []SortParam `json:"sorts"`
}

func (s *SortParam) ToString() string {
	if s.Ascending {
		return s.Field + " ASC"
	}

	return s.Field + " DESC"
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = pagination_const.Page
	}

	if p.Limit <= 0 {
		p.Limit = pagination_const.Limit
	}
}

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

func createDefaultParams() *Param {
	return &Param{
		Page:  Paging{Limit: pagination_const.Limit, Page: pagination_const.Page},
		Sorts: []SortParam{},
	}
}

func ParseSortParam(sortStr string) []SortParam {
	result := make([]SortParam, 0)
	sortElements := strings.Split(sortStr, ",")

	for _, sortElement := range sortElements {
		args := strings.Split(sortElement, ":")
		if len(args) != 2 {
			panic(errors.New("wrong query param for sort: Not enough arguments"))
		} else if strings.ToUpper(args[1]) != paramAscending && strings.ToUpper(args[1]) != paramDescending {
			panic(errors.New("wrong query param for sort: should be asc or desc"))
		} else {
			result = append(result, SortParam{
				Field:     args[0],
				Ascending: strings.ToUpper(args[1]) == paramAscending,
			})
		}
	}

	return result
}

func ParseFilterParam(filterStr string) []ConditionalParam {
	result := make([]ConditionalParam, 0)
	err := json.Unmarshal([]byte(filterStr), &result)

	if err != nil {
		panic(errors.New("cannot parse filter condition: " + err.Error()))
	}

	return result
}

// parse all type of params into the param object
func ParseQueryParams(e gin.Context) (*Param, error) {
	urlValues := e.Request.URL.Query()
	params := createDefaultParams()

	var err error
	for name, values := range urlValues {
		if len(values) == 0 {
			continue
		}

		val := strings.Trim(values[0], " ")

		switch name {
		case paramPage:
			page, err := strconv.Atoi(val)

			if err != nil {
				continue
			}

			params.Page.Page = page

		case paramLimit:
			limit, err := strconv.Atoi(val)

			if err != nil {
				continue
			}

			params.Page.Limit = limit
		case paramSort:
			params.Sorts = ParseSortParam(val)

		case paramFilter:
			params.Filters = ParseFilterParam(val)
		default: // ignore it
		}
	}

	return params, err
}

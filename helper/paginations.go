package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
)

func GeneratePaginationRequest(context *gin.Context) *dto.Pagination {
	// default limit, page & sort parameter
	limit := 5
	page := 0
	sort := "created_at asc"

	var searchs []dto.Search

	query := context.Request.URL.Query()
	fmt.Println(query)

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}

		// check if query parameter key contains dot
		if strings.Contains(key, ".") {
			// split query parameter key by dot
			searchKeys := strings.Split(key, ".")
			// fmt.Println(key)
			// fmt.Println(queryValue)

			// create search object
			// search := request.Search{Column: "id", Action: "equals", Query: "1"}
			search := dto.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &dto.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}

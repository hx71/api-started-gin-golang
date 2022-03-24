package helper

import (
	"srp-golang/app/request"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *request.Pagination {
	// default limit, page & sort parameter
	limit := 5
	page := 1
	sort := "created_at asc"

	var searchs []request.Search

	query := context.Request.URL.Query()

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

			// create search object
			search := request.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &request.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}

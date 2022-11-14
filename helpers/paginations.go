package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	Sort         string      `json:"sort"`
	TotalRows    int64       `json:"total_rows"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
	FromRow      int         `json:"from_row"`
	ToRow        int         `json:"to_row"`
	Rows         interface{} `json:"rows"`
	Searchs      []Search    `json:"searchs"`
}

func GeneratePaginationRequest(context *gin.Context) *Pagination {
	// default limit, page & sort parameter
	limit := 10
	page := 1
	sort := "created_at asc"

	var searchs []Search

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
			search := Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}

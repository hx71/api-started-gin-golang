package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/hx71/api-started-gin-golang/app/role"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"gorm.io/gorm"
)

type roleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(connection *gorm.DB) role.Repository {
	return &roleConnection{
		connection: connection,
	}
}

func (db *roleConnection) Create(model models.Role) error {
	return db.connection.Create(&model).Error
}

func (db *roleConnection) Show(id string) (role models.Role) {
	db.connection.First(&role, "id = ?", id)
	return role
}

func (db *roleConnection) Update(role models.Role) error {
	return db.connection.Updates(&role).Error
}

func (db *roleConnection) Delete(model models.Role) error {
	return db.connection.Delete(&model).Error
}

func (db *roleConnection) Pagination(pagination *response.Pagination) (response.RepositoryResult, int) {
	// Initialize variables
	var records []models.Role
	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	// Calculate offset
	offset := (pagination.Page - 1) * pagination.Limit

	// Generate where query
	var whereConditions []string
	for _, value := range pagination.Searchs {
		column := value.Column
		action := value.Action
		query := value.Query

		switch action {
		case "equals":
			whereConditions = append(whereConditions, fmt.Sprintf("%s = '%s'", column, query))
		case "contains":
			whereConditions = append(whereConditions, fmt.Sprintf("lower(%s) LIKE '%%%s%%'", column, strings.ToLower(query)))
		case "in":
			whereConditions = append(whereConditions, fmt.Sprintf("%s IN ('%s')", column, query))
		}
	}
	// Build the where clause
	where := strings.Join(whereConditions, " AND ")

	// Fetch total rows
	errCount := db.connection.Model(&models.Role{}).Where(where).Count(&totalRows).Error
	if errCount != nil {
		return response.RepositoryResult{Error: errCount}, totalPages
	}

	// Fetch records
	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where(where)
	errFind := find.Find(&records).Error
	if errFind != nil {
		return response.RepositoryResult{Error: errFind}, totalPages
	}

	// Set pagination data
	pagination.Rows = records
	pagination.TotalRows = totalRows

	// Calculate total pages
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	if pagination.Page == 0 {
		// Set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// Calculate from & to row
			fromRow = ((pagination.Page - 1) * pagination.Limit) + 1
			toRow = pagination.Page * pagination.Limit
		}
	}

	if int64(toRow) > totalRows {
		// Set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return response.RepositoryResult{Result: pagination}, totalPages
}

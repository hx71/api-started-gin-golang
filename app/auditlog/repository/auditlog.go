package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/hx71/api-started-gin-golang/app/auditlog"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"gorm.io/gorm"
)

type auditLogConnection struct {
	connection *gorm.DB
}

func NewAuditLogRepository(connection *gorm.DB) auditlog.Repository {
	return &auditLogConnection{connection}
}

func (db *auditLogConnection) Create(model models.AuditLog) error {
	return db.connection.Save(&model).Error
}

func (db *auditLogConnection) Show(id string) (role models.AuditLog) {
	db.connection.Where("id = ?", id).First(&role)
	return role
}

func (db *auditLogConnection) Update(model models.AuditLog) error {
	return db.connection.Updates(&model).Error
}

func (db *auditLogConnection) Delete(model models.AuditLog) error {
	return db.connection.Delete(&model).Error
}

func (db *auditLogConnection) Pagination(pagination *response.Pagination) (response.RepositoryResult, int) {

	var records []models.AuditLog
	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	offset := (pagination.Page - 1) * pagination.Limit

	// get data with limit, offset & order
	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs
	whereEquals := ""
	whereLike := ""
	where := ""

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				if whereEquals == "" {
					whereEquals = fmt.Sprintf(column + " = '" + query + "'")
				} else {
					whereEquals = fmt.Sprintf(whereEquals + " AND " + column + " = '" + query + "'")
				}
				break
			case "contains":
				if whereLike == "" {
					whereLike = fmt.Sprintf("lower("+column+") LIKE  '%%%s%%'", strings.ToLower(query))
				} else {
					whereLike = whereLike + fmt.Sprintf("AND lower("+column+") LIKE  '%%%s%%'", strings.ToLower(query))
				}
				break
			case "in":
				if whereEquals == "" {
					whereEquals = fmt.Sprintf(column + " IN (" + query + ")")
				} else {
					whereEquals = fmt.Sprintf(whereEquals + " AND " + column + " IN (" + query + ")")
				}
				break
			}
		}
		if whereEquals != "" && whereLike != "" {
			where = whereEquals + " AND " + whereLike
		} else {
			where = whereEquals + whereLike
		}
		find = find.Where(where)
		errCount := db.connection.Model(&models.AuditLog{}).Where(where).Count(&totalRows).Error
		if errCount != nil {
			return response.RepositoryResult{Error: errCount}, totalPages
		}
	} else {
		errCount := db.connection.Model(&models.AuditLog{}).Count(&totalRows).Error
		if errCount != nil {
			return response.RepositoryResult{Error: errCount}, totalPages
		}
	}

	find = find.Find(&records)

	// has error find data
	errFind := find.Error
	if errFind != nil {
		return response.RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = records
	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = ((pagination.Page - 1) * pagination.Limit) + 1
			toRow = pagination.Page * pagination.Limit
		}
	}

	if int64(toRow) > totalRows {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return response.RepositoryResult{Result: pagination}, totalPages
}

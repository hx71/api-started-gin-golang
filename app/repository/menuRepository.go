package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(model models.Menu) error
	Show(id string) models.Menus
	Update(model models.Menu) error
	Delete(id string) error
	Pagination(*helpers.Pagination) (RepositoryResult, int)
}

type menuConnection struct {
	connection *gorm.DB
}

//NewMenuRepository is creates a new instance of MenuRepository
func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuConnection{
		connection: db,
	}
}

func (db *menuConnection) Create(model models.Menu) error {
	return db.connection.Save(&model).Error
}

func (db *menuConnection) Show(id string) (role models.Menus) {
	db.connection.Where("id = ?", id).First(&role)
	return role
}

func (db *menuConnection) Update(model models.Menu) error {
	fmt.Println(model.IsActive)
	return db.connection.Updates(&model).Error
}

func (db *menuConnection) Delete(id string) error {
	return db.connection.Where("id = ?", id).Delete(&models.Menus{}).Error
}

func (db *menuConnection) Pagination(pagination *helpers.Pagination) (RepositoryResult, int) {

	var records []models.Menu
	var totalRows int64
	totalPages := 0

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
		errCount := db.connection.Model(&models.Menu{}).Where(where).Count(&totalRows).Error
		if errCount != nil {
			return RepositoryResult{Error: errCount}, totalPages
		}
	} else {
		errCount := db.connection.Model(&models.Menu{}).Count(&totalRows).Error
		if errCount != nil {
			return RepositoryResult{Error: errCount}, totalPages
		}
	}

	find = find.Find(&records)

	// has error find data
	errFind := find.Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = records
	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	if pagination.Page == 0 {
		// set from & to row on first page
		pagination.FromRow = 1
		pagination.ToRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			pagination.FromRow = ((pagination.Page - 1) * pagination.Limit) + 1
			pagination.ToRow = pagination.Page * pagination.Limit
		}
	}

	if int64(pagination.ToRow) > totalRows {
		// set to row with total rows
		pagination.ToRow = int(totalRows)
	}

	return RepositoryResult{Result: pagination}, totalPages
}

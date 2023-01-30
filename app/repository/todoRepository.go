package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(model models.Todo) error
	Show(id string) models.Todos
	Update(model models.Todo) error
	Delete(id string) error
	Pagination(*helpers.Pagination) (RepositoryResult, int)
}

type todoConnection struct {
	connection *gorm.DB
}

//NewTodoRepository is creates a new instance of TodoRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoConnection{
		connection: db,
	}
}

func (db *todoConnection) Create(model models.Todo) error {
	return db.connection.Save(&model).Error
}

func (db *todoConnection) Show(id string) models.Todos {
	var todo models.Todos
	db.connection.Table("todos").
		Select("todos.*, users.id, users.name as name_user").
		Joins("LEFT JOIN users on users.id = todos.user_id").
		Where("todos.id = ?", id).
		Where("todos.created_at is null").
		Scan(&todo)
	return todo
}

func (db *todoConnection) Update(model models.Todo) error {
	return db.connection.Updates(&model).Error
}

func (db *todoConnection) Delete(id string) error {
	var todo models.Todos
	db.connection.Find(&todo, "id = ?", id)
	return db.connection.Delete(&todo).Error
}

func (db *todoConnection) Pagination(pagination *helpers.Pagination) (RepositoryResult, int) {

	var records []models.Todo
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
		errCount := db.connection.Model(&models.Todo{}).Where(where).Count(&totalRows).Error
		if errCount != nil {
			return RepositoryResult{Error: errCount}, totalPages
		}
	} else {
		errCount := db.connection.Model(&models.Todo{}).Count(&totalRows).Error
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

	return RepositoryResult{Result: pagination}, totalPages
}

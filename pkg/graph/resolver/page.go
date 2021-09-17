package resolver

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/domain"
	"github.com/jinzhu/gorm"
)

type direction string

var (
	asc  direction = "asc"
	desc direction = "desc"
)

func pageDB(db *gorm.DB, col string, dir direction, page domain.PaginationInput) (*gorm.DB, error) {
	var limit int
	if page.First == nil {
		limit = 11
	} else {
		limit = *page.First + 1
	}

	if page.After != nil {
		resource1, resource2, err := decodeCursor(*page.After)
		if err != nil {
			return db, err
		}

		if resource2 != nil {
			switch dir {
			case asc:
				db = db.Where(
					fmt.Sprintf("(%s > ?) OR (%s = ? AND id > ?)", col, col),
					resource1.ID,
					resource1.ID, resource2.ID,
				)
			case desc:
				db = db.Where(
					fmt.Sprintf("(%s < ?) OR (%s = ? AND id < ?)", col, col),
					resource1.ID,
					resource1.ID, resource2.ID,
				)
			}
		} else {
			switch dir {
			case asc:
				db = db.Where(fmt.Sprintf("%s > ?", col), resource1.ID)
			case desc:
				db = db.Where(fmt.Sprintf("%s < ?", col), resource1.ID)
			}
		}
	}

	switch dir {
	case asc:
		db = db.Order(fmt.Sprintf("%s IS NULL ASC, id ASC", col))
	case desc:
		db = db.Order(fmt.Sprintf("%s DESC, id DESC", col))
	}

	return db.Limit(limit), nil
}

type cursorResource struct {
	Name string
	ID   int
}

func createCursor(first cursorResource, second *cursorResource) string {
	var cursor []byte
	if second != nil {
		cursor = []byte(fmt.Sprintf("%s:%d:%s:%d", first.Name, first.ID, second.Name, second.ID))
	} else {
		cursor = []byte(fmt.Sprintf("%s:%d", first.Name, first.ID))
	}

	return base64.StdEncoding.EncodeToString(cursor)
}

func decodeCursor(cursor string) (cursorResource, *cursorResource, error) {
	bytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return cursorResource{}, nil, err
	}

	vals := strings.Split(string(bytes), ":")

	switch len(vals) {
	case 2:
		id, err := strconv.Atoi(vals[1])
		if err != nil {
			return cursorResource{}, nil, errors.New("invalid_cursor")
		}

		return cursorResource{Name: vals[0], ID: id}, nil, nil
	case 4:
		id, err := strconv.Atoi(vals[1])
		if err != nil {
			return cursorResource{}, nil, errors.New("invalid_cursor")
		}

		id2, err := strconv.Atoi(vals[3])
		if err != nil {
			return cursorResource{}, nil, errors.New("invalid_cursor")
		}

		return cursorResource{
				Name: vals[0],
				ID:   id,
			}, &cursorResource{
				Name: vals[2],
				ID:   id2,
			}, nil
	default:
		return cursorResource{}, nil, errors.New("invalid_cursor")
	}
}

func convertToConnection(tasks []*domain.Task, orderBy domain.TaskOrderFields, page domain.PaginationInput) *domain.TaskConnection {
	if len(tasks) == 0 {
		return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}
	}

	pageInfo := domain.PageInfo{}
	if page.First != nil {
		if len(tasks) >= *page.First+1 {
			pageInfo.HasNextPage = true
			tasks = tasks[:len(tasks)-1]
		}
	}

	switch orderBy {
	case domain.TaskOrderFieldsLatest:
		taskEdges := make([]*domain.TaskEdge, len(tasks))

		for i, task := range tasks {
			cursor := createCursor(
				cursorResource{Name: "task", ID: task.ID},
				nil,
			)
			taskEdges[i] = &domain.TaskEdge{
				Cursor: cursor,
				Node:   task,
			}
		}

		pageInfo.EndCursor = taskEdges[len(taskEdges)-1].Cursor

		return &domain.TaskConnection{PageInfo: &pageInfo, Edges: taskEdges}
	case domain.TaskOrderFieldsDue:
		taskEdges := make([]*domain.TaskEdge, 0, len(tasks))

		for _, task := range tasks {
			if task.Due == "" {
				pageInfo.HasNextPage = false
				return &domain.TaskConnection{PageInfo: &pageInfo, Edges: taskEdges}
			}

			cursor := createCursor(
				cursorResource{Name: "task", ID: task.ID},
				&cursorResource{Name: "due", ID: task.ID},
			)

			taskEdges = append(taskEdges, &domain.TaskEdge{
				Cursor: cursor,
				Node:   task,
			})
		}

		pageInfo.EndCursor = taskEdges[len(taskEdges)-1].Cursor

		return &domain.TaskConnection{PageInfo: &pageInfo, Edges: taskEdges}
	}

	return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}
}

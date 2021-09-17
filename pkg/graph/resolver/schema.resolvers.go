package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/domain"
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	db := r.DB

	user := domain.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := db.Create(&user).Error; err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error) {
	db := r.DB

	_, sequentialID := fromGlobalID(input.ID)

	var user domain.User
	if err := db.Where("id = ?", sequentialID).First(&user).Error; err != nil {
		return &domain.User{}, err
	}

	params := map[string]interface{}{}
	if input.Name != nil {
		params["name"] = *input.Name
	}
	if input.Email != nil {
		params["email"] = *input.Email
	}

	if err := db.Model(&user).Updates(params).Error; err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input domain.DeleteUserInput) (*domain.User, error) {
	db := r.DB

	_, sequentialID := fromGlobalID(input.ID)

	var user domain.User
	if err := db.Where("id = ?", sequentialID).First(&user).Error; err != nil {
		return &domain.User{}, err
	}

	if err := db.Delete(&user).Error; err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, input domain.CreateTaskInput) (*domain.Task, error) {
	db := r.DB

	_, sequentialID := fromGlobalID(input.UserID)

	// タイムゾーン指定
	var timeZoneJST = time.FixedZone("Asia/Tokyo", 9*60*60)
	time.Local = timeZoneJST
	time.LoadLocation("Asia/Tokyo")

	task := domain.Task{
		Title:  input.Title,
		UserID: sequentialID,
		Due:    input.Due,
	}

	if input.Notes != nil {
		task.Notes = *input.Notes
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := db.Create(&task).Error; err != nil {
		return &domain.Task{}, err
	}

	return &task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input domain.UpdateTaskInput) (*domain.Task, error) {
	db := r.DB

	_, sequentialID := fromGlobalID(input.ID)

	var task domain.Task
	if err := db.Where("id = ?", sequentialID).First(&task).Error; err != nil {
		return &domain.Task{}, err
	}

	params := map[string]interface{}{}
	if input.Title != nil {
		params["title"] = *input.Title
	}
	if input.Notes != nil {
		params["notes"] = *input.Notes
	}
	if input.Completed != nil {
		params["completed"] = *input.Completed
	}
	if input.Due == nil {
		params["due"] = nil
	} else {
		params["due"] = *input.Due
	}

	if err := db.Model(&task).Updates(params).Error; err != nil {
		return &domain.Task{}, err
	}

	return &task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input domain.DeleteTaskInput) (*domain.Task, error) {
	db := r.DB

	_, sequentialID := fromGlobalID(input.ID)

	var task domain.Task
	if err := db.Where("id = ?", sequentialID).First(&task).Error; err != nil {
		return &domain.Task{}, err
	}

	if err := db.Delete(&task).Error; err != nil {
		return &domain.Task{}, err
	}

	return &task, nil
}

func (r *queryResolver) Users(ctx context.Context, orderBy domain.UserOrderFields, page domain.PaginationInput) (*domain.UserConnection, error) {
	db := r.DB

	var users []*domain.User
	db.Limit(*page.First).Find(&users)

	pageInfo := domain.PageInfo{}
	if page.First != nil {
		if len(users) >= *page.First+1 {
			pageInfo.HasNextPage = true
			users = users[:len(users)-1]
		}
	}

	userEdges := make([]*domain.UserEdge, len(users))

	for i, user := range users {
		cursor := createCursor(
			cursorResource{Name: "user", ID: user.ID},
			nil,
		)
		userEdges[i] = &domain.UserEdge{
			Cursor: cursor,
			Node:   user,
		}
	}

	if len(userEdges) > 0 {
		pageInfo.EndCursor = userEdges[len(userEdges)-1].Cursor
	}

	return &domain.UserConnection{PageInfo: &pageInfo, Edges: userEdges}, nil
}

func (r *queryResolver) Tasks(ctx context.Context, input domain.TasksInput, orderBy domain.TaskOrderFields, page domain.PaginationInput) (*domain.TaskConnection, error) {
	db := r.DB

	if input.Completed != nil {
		db = db.Where("completed = ?", *input.Completed)
	}

	var err error

	switch orderBy {
	case domain.TaskOrderFieldsLatest:
		db, err = pageDB(db, "id", desc, page)
		if err != nil {
			return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}, err
		}

		var tasks []*domain.Task
		if err := db.Find(&tasks).Error; err != nil {
			return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}, err
		}

		return convertToConnection(tasks, orderBy, page), nil
	case domain.TaskOrderFieldsDue:
		db, err = pageDB(db, "UNIX_TIMESTAMP(due)", asc, page)
		if err != nil {
			return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}, err
		}

		var tasks []*domain.Task
		if err := db.Find(&tasks).Error; err != nil {
			return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}, err
		}

		return convertToConnection(tasks, orderBy, page), nil
	default:
		return &domain.TaskConnection{PageInfo: &domain.PageInfo{}}, errors.New("invalid order by")
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

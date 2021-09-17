package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/domain"
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/generated"
)

func (r *userResolver) ID(ctx context.Context, obj *domain.User) (string, error) {
	if obj == nil {
		return "", nil
	}

	fmt.Println(ctx)

	globalID := toGlobalID("user", obj.ID)

	return globalID, nil
}

func (r *userResolver) Tasks(ctx context.Context, obj *domain.User) ([]*domain.Task, error) {
	db := r.DB

	var tasks []*domain.Task
	db.Where("user_id = ?", obj.ID).Find(&tasks)

	return tasks, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

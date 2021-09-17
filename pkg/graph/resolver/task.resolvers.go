package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/domain"
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/generated"
)

func (r *taskResolver) ID(ctx context.Context, obj *domain.Task) (string, error) {
	// 各GraphQL Schemaのフィールド名で関数を作ると、各フィールド用のリゾルバが作成できる。
	if obj == nil {
		return "", nil
	}

	globalID := toGlobalID("task", obj.ID)

	return globalID, nil
}

func (r *taskResolver) OwnedBy(ctx context.Context, obj *domain.Task) (*domain.User, error) {
	db := r.DB

	user := &domain.User{}
	db.Where("id = ?", obj.UserID).First(&user)

	return user, nil
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *taskResolver) Due(ctx context.Context, obj *domain.Task) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fukunokaze/mfs_backend/APIGateway/graph/generated"
	"github.com/fukunokaze/mfs_backend/APIGateway/graph/model"
)

func (r *mutationResolver) Authenticate(ctx context.Context, input model.Login) (*model.User, error) {
	return &model.User{Name: "OK", ID: "1"}, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

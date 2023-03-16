package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"mfs_backend/LoginService/loginpb"

	"github.com/fukunokaze/mfs_backend/APIGateway/graph/generated"
	"github.com/fukunokaze/mfs_backend/APIGateway/graph/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (r *mutationResolver) Authenticate(ctx context.Context, input model.Login) (*model.User, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := loginpb.NewInventoryClient(conn)

	client.VerifyLogin()

	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
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

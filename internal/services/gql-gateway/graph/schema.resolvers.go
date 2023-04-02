package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"fmt"
	"jikkaem/internal/services/gql-gateway/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Fancam is the resolver for the fancam field.
func (r *queryResolver) Fancam(ctx context.Context) ([]*model.Fancam, error) {
	panic(fmt.Errorf("not implemented: Fancam - fancam"))
}

// Artist is the resolver for the artist field.
func (r *queryResolver) Artist(ctx context.Context) ([]*model.Artist, error) {
	panic(fmt.Errorf("not implemented: Artist - artist"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, input model.SingleUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

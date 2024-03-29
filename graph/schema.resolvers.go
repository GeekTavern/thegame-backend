package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"
	"thegame-backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreateStoryPoint is the resolver for the createStoryPoint field.
func (r *mutationResolver) CreateStoryPoint(ctx context.Context, text string, image string) (string, error) {
	panic(fmt.Errorf("not implemented: CreateStoryPoint - createStoryPoint"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Storypoint is the resolver for the storypoint field.
func (r *queryResolver) Storypoint(ctx context.Context) (*model.StoryPoint, error) {
	var storypoint *model.StoryPoint
	dummyStoryPoint := model.StoryPoint{
		ID:   "1",
		Text: "This is the first story point!",
	}
	storypoint = &dummyStoryPoint
	return storypoint, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

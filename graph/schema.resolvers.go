package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/jlamb1/go-graphql/graph/generated"
	"github.com/jlamb1/go-graphql/graph/model"
	"github.com/jlamb1/go-graphql/repository"
)

func (r *mutationResolver) CreateBlogPost(ctx context.Context, input model.NewBlogPost) (*model.BlogPost, error) {
	rand.Seed(time.Now().UnixNano())
	blogPost := &model.BlogPost{
		ID:     strconv.Itoa(rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "user " + input.UserID},
		Body:   input.Body,
	}
	blogPostRepo.Save(blogPost)
	return blogPost, nil
}

func (r *queryResolver) BlogPosts(ctx context.Context) ([]*model.BlogPost, error) {
	return blogPostRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var blogPostRepo repository.BlogRepository = repository.New()

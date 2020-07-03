package graph

import "github.com/jlamb1/go-graphql/graph/model"

//go:generate go run runner/runner.go

//Resolver default setup
type Resolver struct {
	videos []*model.Video
}

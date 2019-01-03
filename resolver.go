package graphql_tutorial

import (
	"context"
)

var dummyTodos = []Todo{
	Todo{
		Text: "hello",
		ID: "1",
		Done: false,
		User: "jun",
	},
	Todo{
		Text: "hello1",
		ID: "2",
		Done: false,
		User: "yun",
	},
}

var dummyUsers = []User{
	User{
		ID: "1",
		Name: "jun",
		Age: 123,
	},
	User{
		ID: "1",
		Name: "jun",
		Age: 123,
	},
}

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return dummyTodos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	return dummyUsers, nil
}

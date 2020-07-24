package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/CKjapan/go-graphql-practice/server/dao"
	"github.com/CKjapan/go-graphql-practice/server/graph/generated"
	models "github.com/CKjapan/go-graphql-practice/server/models/defined-model"
	"github.com/CKjapan/go-graphql-practice/server/models/generated-model"
	"github.com/CKjapan/go-graphql-practice/server/util"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	id := util.CreateUniqueID()
	err := dao.NewUserDao(r.DB).InsertOne(&dao.User{
		ID:   id,
		Name: input.Name,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

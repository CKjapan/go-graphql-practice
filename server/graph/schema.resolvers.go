package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/CKjapan/go-graphql-practice/server/dao"
	"github.com/CKjapan/go-graphql-practice/server/graph/generated"
	models "github.com/CKjapan/go-graphql-practice/server/models/defined-model"
	"github.com/CKjapan/go-graphql-practice/server/models/generated-model"
	"github.com/CKjapan/go-graphql-practice/server/util"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	id := util.CreateUniqueID()
	err := dao.NewTodoDao(r.DB).InsertOne(&dao.Todo{
		ID:     id,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		return "", err
	}
	return id, nil
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
	// log.Printf("[mutationResolver.DeleteTodo] id: %s", input)
	// todo, err := dao.NewTodoDao(r.DB).DeleteOne(input)
	// if err != nil {
	// 	return nil, err
	// }
	// if todo == nil {
	// 	return nil, errors.New("not found")
	// }
	// return &models.Todo{
	// 	ID:   todo.ID,
	// 	Text: todo.Text,
	// 	Done: todo.Done,
	// }, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Println("[queryResolver.Todos]")
	todos, err := dao.NewTodoDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	log.Printf("[queryResolver.Todo] id: %s", id)
	todo, err := dao.NewTodoDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("not found")
	}
	return &models.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	log.Println("[queryResolver.Users]")
	users, err := dao.NewUserDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.User
	for _, user := range users {
		results = append(results, &models.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	user, err := dao.NewUserDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("not found")
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	log.Printf("[todoResolver.User] id: %#v", obj)
	user, err := dao.NewUserDao(r.DB).FindByTodoID(obj.ID)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	log.Println("[userResolver.Todos]")
	todos, err := dao.NewTodoDao(r.DB).FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
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

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/touchps/hackernews/config"
	"github.com/touchps/hackernews/dal"
	"github.com/touchps/hackernews/graph/generated"
	"github.com/touchps/hackernews/graph/model"
	"github.com/touchps/hackernews/types"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	// panic(fmt.Errorf("not implemented"))
	var link types.Link
	// var user model.User
	link.Address = input.Address
	link.Title = input.Title
	// user.Name = "test"
	// link.User = &user
	mysqldal := dal.MySQLDAL{
		SqlDB: config.GetSqlDB(),
	}
	linkID, err := mysqldal.SaveLink(link)
	if err != nil {
		return nil, err
	}
	return &model.Link{
		ID:      strconv.FormatInt(linkID, 10),
		Title:   link.Title,
		Address: link.Address,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// panic(fmt.Errorf("not implemented"))
	mysqldal := dal.MySQLDAL{
		SqlDB: config.GetSqlDB(),
	}
	var links []*model.Link
	linksData, err := mysqldal.GetLinks()
	if err != nil {
		return nil, err
	}

	for _, link := range linksData {
		links = append(links, &model.Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
		})
	}
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

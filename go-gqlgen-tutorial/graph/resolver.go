package graph

import (
	"context"
	"strconv"

	"github.com/leehom/go-gqlgen-tutorial/graph/model"
	"github.com/leehom/go-gqlgen-tutorial/internal/links"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
// 	var links []*model.Link
// 	links = append(links, &model.Link{Title: "our dummy link", Address: "https://address.org", User: &model.User{Name: "admin"}})
// 	return links, nil
// }

// func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
// 	var link model.Link
// 	var user model.User
// 	link.Address = input.Address
// 	link.Title = input.Title
// 	user.Name = "test"
// 	link.User = &user
// 	return &link, nil
// }

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkId := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkId, 10), Title: link.Title, Address: link.Address}, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}
	return resultLinks, nil
}

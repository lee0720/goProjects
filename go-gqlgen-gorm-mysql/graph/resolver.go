package graph

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/leehom/go-gqlgen-gorm-mysql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func mapItemsFromInput(itemsInput []*model.ItemInput) []model.Item {
	var items []model.Item
	for _, itemInput := range itemsInput {
		items = append(items, model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantify:    itemInput.Quantity,
		})
	}
	return items
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {

	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        input.Items,
	}
	r.DB.Create(&order)
	return &order, nil
}

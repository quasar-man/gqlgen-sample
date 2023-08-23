package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"context"
	"gqlgen-sample/db"
	"gqlgen-sample/graph/model"
)

type Resolver struct{}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	db, err := db.ConnectDB()
	if err != nil {
			return nil, err
	}

	user := model.User{}
	err = db.First(&user, id).Error
	if err != nil {
			return nil, err
	}
	return &user, nil
}

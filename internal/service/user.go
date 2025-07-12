package service

import (
	"github.com/zakkbob/mxguard/internal/model"
	"context"
)

type UserRepository interface {
	CreateUser(context.Context, string, bool) (model.User, error)
}

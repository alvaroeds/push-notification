package repository

import "context"

type Repository interface {
	GetTokensByUser(ctx context.Context, user string) ([]string, error)
	UpdateTokensByUser(ctx context.Context, user string, tokens []string) error
}

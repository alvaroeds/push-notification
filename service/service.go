package service

import "context"

type Service interface {
	GetTokensByUsers(ctx context.Context, usersID []string) (map[string][]string, error)
	UpdateTokensByUser(ctx context.Context, tokens []string, user string) error
}

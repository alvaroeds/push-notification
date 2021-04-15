package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"dinamo.app/push_notification/repository"
	"log"
)

type firestoreService struct {
	repo repository.Repository
}

func (s *firestoreService) GetTokensByUsers(ctx context.Context, usersID []string) (map[string][]string, error) {
	ut := map[string][]string{}
	for _, user := range usersID {
		tokens, err := s.repo.GetTokensByUser(ctx, user)
		if err != nil {
			log.Println(err)
			continue
		} else if len(tokens) != 0 {
			ut[user] = tokens
		}
	}

	return ut, nil
}

func (s *firestoreService) UpdateTokensByUser(ctx context.Context, tokens []string, user string) error {
	err := s.repo.UpdateTokensByUser(ctx, user, tokens)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// NewService returns a new postgres repository.
func NewService(db *firestore.Client) Service {
	return &firestoreService{
		repo: repository.NewRepository(db),
	}
}

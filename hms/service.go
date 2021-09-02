package hms

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type firestoreService struct {
	repo Repository
}

func (s *firestoreService) GetTokensByUsers(ctx context.Context, usersID []string) ([]string, error) {
	ut := []string{}
	for _, user := range usersID {
		tokens, err := s.repo.GetTokensByUser(ctx, user)
		if err != nil {
			log.Println(err)
			continue
		}

		ut = append(ut, tokens...)
	}

	return ut, nil
}

// NewService returns a new postgres repository.
func NewService(db *firestore.Client) Service {
	return &firestoreService{
		repo: newRepository(db),
	}
}

type Service interface {
	GetTokensByUsers(ctx context.Context, usersID []string) ([]string, error)
	PostNotificationPush(data []byte) error
}

func (s *firestoreService) PostNotificationPush(data []byte) error {
	return s.repo.postNotificationPush(data)
}

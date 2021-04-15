package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type FirebaseRepository struct {
	db *firestore.Client
}

func (f *FirebaseRepository) GetTokensByUser(ctx context.Context, user string) ([]string, error) {
	type UserTokens struct {
		Tokens []string `firestore:"tokens,omitempty"`
	}

	doc, err := f.db.Collection("user-token-devices").Doc(user).Get(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ut := UserTokens{}
	err = doc.DataTo(&ut)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ut.Tokens, nil
}

func (f *FirebaseRepository) UpdateTokensByUser(ctx context.Context, user string, tokens []string) error {

	_, err := f.db.Collection("user-token-devices").Doc(user).Update(ctx, []firestore.Update{
		{
			Path:  "tokens",
			Value: tokens,
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// NewRepository returns a new postgres repository.
func NewRepository(db *firestore.Client) Repository {
	return &FirebaseRepository{
		db: db,
	}
}

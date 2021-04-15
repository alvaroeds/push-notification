package repository

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./dinamo-fa84e-firebase-adminsdk-o6dr2-390cdf4656.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)

	}

	clientF, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

	fr := NewRepository(clientF)
	tokens, err := fr.GetTokensByUser(ctx, "9zpyJF0H3XYA6IIUc11CNSTnkII2")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tokens)
}

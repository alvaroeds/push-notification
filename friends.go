package push_notification

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"fmt"
	"strings"
)

func FriendsReciver(ctx context.Context, e FirestoreEvent) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection := pathParts[0]
	uidR := pathParts[1]
	uidS := pathParts[3]
	doc := strings.Join(pathParts[1:], "/")

	fmt.Println(e)
	fmt.Println(fullPath)
	fmt.Println(pathParts)
	fmt.Println(collection)
	fmt.Println(uidR)
	fmt.Println(uidS)
	fmt.Println(doc)

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	fmt.Println("//////")
	fmt.Println(*meta.Resource)



	return nil
}
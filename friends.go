package push_notification

import (
	"context"
	"fmt"
	"strings"
)

func FriendsReciver(ctx context.Context, e FirestoreEvent) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection := pathParts[0]
	doc := strings.Join(pathParts[1:], "/")

	fmt.Println(e)
	fmt.Println(collection)
	fmt.Println(doc)

	return nil
}
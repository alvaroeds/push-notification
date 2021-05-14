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
	fmt.Println(ctx.Value("auth"))

	return nil
}
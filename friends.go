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

	uidReciver := ctx.Value("uidReciver")
	//idR, _ := uidReciver.(string)
	uidSender := ctx.Value("uidSender")
	//idS, _ := uidSender.(string)
	//if !ok {
	//	return er
	//}
	//return id


	fmt.Println(e)
	fmt.Println(fullPath)
	fmt.Println(collection)
	fmt.Println(doc)
	fmt.Println("///")
	fmt.Println(uidReciver)
	fmt.Println(uidSender)

	return nil
}
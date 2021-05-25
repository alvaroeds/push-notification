package push_notification

import (
	"context"
	"firebase.google.com/go/messaging"
	"log"
	"strings"
)

const (
	SOLICITUD_ADD_CONTACT = "SOLICITUD_ADD_CONTACT"
)

func FriendsReciver(ctx context.Context, e FirestoreEvent) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	//collection := pathParts[0]
	uidR := pathParts[1]
	uidS := pathParts[3]
	//doc := strings.Join(pathParts[1:], "/")

	name := e.Value.Fields.Name.StringValue
	if name == "" {
		name = e.Value.Fields.Nickname.StringValue
	}
	body := string(name + " quiere ser parte de tu red Dinamo")

	data := DataNotification{
		UsersID: []string{
			uidR,
		},
		Message: &messaging.MulticastMessage{
			Tokens: nil,
			Data: map[string]string{
				"type":      SOLICITUD_ADD_CONTACT,
				"id_sender": uidS,
				"avatar":    e.Value.Fields.Photo.StringValue,
			},
			Notification: &messaging.Notification{
				Title:    "Nueva Solicitud de amistad",
				Body:     body,
				ImageURL: e.Value.Fields.Photo.StringValue,
			},
		},
	}

	err := SendPushNotificaction(ctx, &data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

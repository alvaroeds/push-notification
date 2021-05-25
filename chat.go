package push_notification

import (
	"context"
	"errors"
	"firebase.google.com/go/messaging"
	"log"
	"strings"
)

const (
	MESSAGE_CHAT    = "MESSAGE_CHAT"
	IMAGE           = "IMAGE"
	IMAGE_OFFLINE   = "IMAGE_OFFLINE"
	STICKER         = "STICKER"
	STICKER_OFFLINE = "STICKER_OFFLINE"
)

func MessageReciver(ctx context.Context, e FirestoreEvent) error {
	if e.Value.Fields.IDTo.StringValue == "" {
		return errors.New("SE REQUIERE UN ID MINIMO PARA ENVIAR NOTIFICACION")
	}

	type_ := e.Value.Fields.Type_.StringValue
	switch type_ {
	case IMAGE, IMAGE_OFFLINE:
		if e.Value.Fields.Content.StringValue == "" {
			e.Value.Fields.Content.StringValue = "📷 Te envió una imagen."
		}
	case STICKER, STICKER_OFFLINE:
		e.Value.Fields.Content.StringValue = "🌚 Te envió un sticker."
	}

	if len(e.Value.Fields.Content.StringValue) >= 50 {
		e.Value.Fields.Content.StringValue = e.Value.Fields.Content.StringValue[0:50]
	}

	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")

	data := DataNotification{
		UsersID: []string{
			e.Value.Fields.IDTo.StringValue,
		},
		Message: &messaging.MulticastMessage{
			Tokens: nil,
			Data: map[string]string{
				"type":        MESSAGE_CHAT,
				"sub_type":    e.Value.Fields.Type_.StringValue,
				"uid_contact": e.Value.Fields.IDFrom.StringValue,
				"avatar":      e.Value.Fields.AvatarFrom.StringValue,
				"id_message":  pathParts[3],
			},
			Notification: &messaging.Notification{
				Title:    e.Value.Fields.NameFrom.StringValue,
				Body:     e.Value.Fields.Content.StringValue,
				ImageURL: e.Value.Fields.Image.StringValue,
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

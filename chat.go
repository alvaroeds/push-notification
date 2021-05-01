package push_notification

import (
	"context"
	"errors"
	"firebase.google.com/go/messaging"
	"log"
	"strconv"
)

const (
	MESSAGE_TEXT = "MESSAGE_TEXT"
	MESSAGE_IMAGE = "MESSAGE_IMAGE"
)

func MessageReciver(ctx context.Context, e FirestoreEvent) error {
	if e.Value.Fields.IDTo.StringValue == "" {
		return errors.New("SE REQUIERE UN ID MINIMO PARA ENVIAR NOTIFICACION")
	}

	image := ""
	dataType := MESSAGE_TEXT
	type_, _ := strconv.Atoi(e.Value.Fields.Type_.IntegerValue)
	if type_ == 2 {
		dataType = MESSAGE_IMAGE
		image = e.Value.Fields.Content.StringValue
		e.Value.Fields.Content.StringValue = "Te envió una imagen."
	}

	data := DataNotification{
		UsersID: []string{
			e.Value.Fields.IDTo.StringValue,
		},
		Message: &messaging.MulticastMessage{
			Tokens: nil,
			Data: map[string]string{
				"type":       dataType,
				"id-contact": e.Value.Fields.IDFrom.StringValue,
			},
			Notification: &messaging.Notification{
				Title:    e.Value.Fields.NameFrom.StringValue,
				Body:     e.Value.Fields.Content.StringValue,
				ImageURL: image,
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

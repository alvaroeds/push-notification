package push_notification

import "firebase.google.com/go/messaging"

type DataNotification struct {
usersID []string                    `json:"users_id"`
message *messaging.MulticastMessage `json:"message"`
}

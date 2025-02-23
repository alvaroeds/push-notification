package push_notification

import (
	"firebase.google.com/go/messaging"
	"time"
)

type DataNotification struct {
	UsersID []string                    `json:"users_id"`
	Message *messaging.MulticastMessage `json:"message"`
}

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log an interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     MyData    `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

type MyData struct {
	AvatarFrom tipos `json:"avatarFrom"`
	Content    tipos `json:"content"`
	Image      tipos `json:"image"`
	IDFrom     tipos `json:"idFrom"`
	IDTo       tipos `json:"idTo"`
	NameFrom   tipos `json:"nameFrom"`
	Type_      tipos `json:"type"`

	//Friends
	Name     tipos `json:"name"`
	Photo    tipos `json:"photoUrlMin"`
	Nickname tipos `json:"nickname"`
}

type tipos struct {
	StringValue  string `json:"stringValue"`
	IntegerValue string `json:"integerValue"`
}

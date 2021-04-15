package push_notification

import (
	"context"
	"dinamo.app/push_notification/response"
	"dinamo.app/push_notification/service"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = "dinamo-fa84e"
var s service.Service
var clientP *messaging.Client
var wg sync.WaitGroup

func init() {
	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: projectID}
	//opt := option.WithCredentialsFile("./dinamo-fa84e-firebase-adminsdk-o6dr2-390cdf4656.json")

	// Use context.Background() because the app/clientF should persist across
	// invocations.
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	//firestore
	f, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
	s = service.NewService(f)

	// Obtain a messaging.Client from the App.
	clientP, err = app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
}

func SendPushNotificaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		_ = response.HTTPError(w, r, http.StatusBadRequest, "Metodo POST requerido")
		return
	}
	ctx := r.Context()
	data := struct {
		UsersID []string                    `json:"users_id"`
		Message *messaging.MulticastMessage `json:"message"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if len(data.UsersID) == 0 {
		err = errors.New("No procede la solicitud sin un UserID")
		log.Println(err)
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//toker organized by users
	Utokens, err := s.GetTokensByUsers(ctx, data.UsersID)
	if err != nil {
		log.Println(err)
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	wg.Add(len(Utokens))
	for id, ut := range Utokens {
		tokens := []string{}
		for _, t := range ut {
			if !duplicatedConfirm(tokens, t) {
				tokens = append(tokens, t)
			}
		}
		go confirmAndUpdateSendMessage(tokens, ut, id)
		data.Message.Tokens = append(data.Message.Tokens, tokens...)
	}

	_, err = clientP.SendMulticast(ctx, data.Message)
	if err != nil {
		log.Println(err)
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	wg.Wait()
	_ = response.HTTPError(w, r, http.StatusOK, "OK")
	return
}

func confirmAndUpdateSendMessage(tokens, tokens_ []string, user string) {
	defer wg.Done()
	Vtokens := []string{}

	batch_, _ := clientP.SendMulticastDryRun(context.Background(), &messaging.MulticastMessage{Tokens: tokens})
	for i, rsp := range batch_.Responses {
		if !rsp.Success {
			fmt.Println(rsp.MessageID)
			continue
		}
		Vtokens = append(Vtokens, tokens[i])
	}

	if len(tokens_) != len(Vtokens) || len(tokens) != len(tokens_) {
		err := s.UpdateTokensByUser(context.Background(), Vtokens, user)
		if err != nil {
			log.Println(err)
		}
	}
}

func duplicatedConfirm(tokens []string, token string) bool {
	r := false
	for _, t := range tokens {
		if t == token {
			r = true
		}
	}

	return r
}

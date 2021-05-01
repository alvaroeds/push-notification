package push_notification

import (
	"dinamo.app/push_notification/response"
	"encoding/json"
	"log"
	"net/http"
)

func HandlerSendPushNotificaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		_ = response.HTTPError(w, r, http.StatusBadRequest, "Metodo POST requerido")
		return
	}
	ctx := r.Context()
	data := DataNotification{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	err = SendPushNotificaction(ctx, &data)
	if err != nil {
		log.Println(err.Error())

		return
	}

	_ = response.JSON(w, r, http.StatusBadRequest, "OK")
}

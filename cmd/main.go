package main

import (
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
	"net/http"
	"push_notification"
	"push_notification/response"
)

func main()  {
	http.HandleFunc("/prueba", push_notification.SendPushNotificaction)
	http.HandleFunc("/body", ResponseBody)

	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}

func ResponseBody(w http.ResponseWriter, r *http.Request) {

	data := struct {
		UsersID []string                    `json:"users_id"`
		Message *messaging.MulticastMessage `json:"message"`
	}{}

	message := &messaging.MulticastMessage{
		Tokens: nil,
		Data:   nil,
		Notification: &messaging.Notification{
			Title:    "PRUEBA DESDE GOLANG 2",
			Body:     "ESTE ES EL BODY DE LA PRUEBA 2",
			ImageURL: "https://i.blogs.es/5efe2c/cap_001/450_1000.jpg",
		},
		Android: &messaging.AndroidConfig{
			CollapseKey:           "asd_:",
			Priority:              "",
			TTL:                   nil,
			RestrictedPackageName: "",
			Data:                  nil,
			Notification: &messaging.AndroidNotification{
				Title: "notificacion android",
				Body:  "body notificacion 1",
				Icon:  "",
				Color: "#9B2BE5",
				Sound: "",
				//remplaza la notificacion por la siguiente con el mismo tag
				Tag:                   "",
				ClickAction:           "",
				BodyLocKey:            "",
				BodyLocArgs:           nil,
				TitleLocKey:           "",
				TitleLocArgs:          nil,
				ChannelID:             "",
				ImageURL:              "https://i.blogs.es/b6d70c/rick-y-morty/1366_521.jpeg",
				Ticker:                "",
				Sticky:                false,
				EventTimestamp:        nil,
				LocalOnly:             false,
				Priority:              0,
				VibrateTimingMillis:   nil,
				DefaultVibrateTimings: false,
				DefaultSound:          false,
				LightSettings:         nil,
				DefaultLightSettings:  false,
				Visibility:            0,
				//NotificationCount:     &[]int{1}[0],
			},
			FCMOptions: nil,
		},
		Webpush: nil,
		APNS:    nil,
	}

	data.Message = message
	data.UsersID = []string{
		"DGdlVlQuGUVjgW5xyysx9mH8P2v2",
		"9zpyJF0H3XYA6IIUc11CNSTnkII2",
	}


	_ = response.JSON(w, r, http.StatusOK, data)
	return
}
package main

import (
	"context"
	"mime"
	"os"
	"path/filepath"

	"github.com/Simplou/meta/whatsapp"
)

func main() {
	ctx := context.Background()
	type enviroment struct {
		accessToken string
		senderID    string
		phoneNumber string
	}
	getEnv := func() *enviroment {
		return &enviroment{
			senderID:    os.Getenv("SENDER_ID"),
			accessToken: os.Getenv("ACCESS_TOKEN"),
			phoneNumber: os.Getenv("PHONE_NUMBER"),
		}
	}
	env := getEnv()
	whats := whatsapp.New(ctx, "v19.0", env.accessToken, env.senderID)
	filePath := "./cmd/build.png"
	ext := filepath.Ext(filePath)
	imageID, err := whatsapp.GenerateMediaID(whats, filePath, mime.TypeByExtension(ext))
	if err != nil {
		panic(err)
	}
	err = whatsapp.SendMessage(
		whats,
		whatsapp.WhatsappBody[whatsapp.Image]{
			JSON: whatsapp.Image{
				Header: whatsapp.Header{
					MessagingProduct: "whatsapp",
					RecipientType:    "individual",
					To:               env.phoneNumber,
					Type:             "image",
				},
				Image: whatsapp.Media{
					Id: imageID,
				},
			},
		})
	if err != nil {
		panic(err)
	}
	println("mensagem enviada com sucesso")
}

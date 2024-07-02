package whatsapp

import (
	"encoding/json"
	"io"

	"github.com/Simplou/goxios"
)

type Header struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to"`
	Type             string `json:"type"`
}

type TextBody struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type Text struct {
	Header
	Text TextBody `json:"text"`
}

// Media types
type (
	Media struct {
		Id       string `json:"id"`                 // your media id
		Filename string `json:"filename,omitempty"` //your document filename
	}
	Audio struct {
		Header
		Audio Media `json:"audio"`
	}
	Document struct {
		Header
		Document Media `json:"document"`
	}
	Image struct {
		Header
		Image Media `json:"image"`
	}
	Video struct {
		Header
		Video Media `json:"video"`
	}
	Sticker struct {
		Header
		Sticker Media `json:"sticker"`
	}
)

type WhatsappBody[MessageType Text | Audio | Document | Image | Video | Sticker] struct {
	JSON MessageType
}

func SendMessage[MessageType Text | Audio | Document | Image | Video | Sticker](whatsapp *client, body WhatsappBody[MessageType]) error {
	b, err := json.Marshal(body.JSON)
	if err != nil {
		return err
	}

	url := whatsapp.Endpoint(whatsapp.senderID + "/messages")
	res, err := whatsapp.Post(url, &goxios.RequestOpts{
		Body:    whatsapp.Reader(b),
		Headers: whatsapp.Headers("application/json"),
	})
	if err != nil {
		return err
	}
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := validateStatus(res, b); err != nil {
		return err
	}
	if err := res.Body.Close(); err != nil {
		return err
	}
	return nil
}

package whatsapp

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/Simplou/goxios"
)

func GenerateMediaID(whatsapp *client, filePath, mimeType string) (string, error) {
	var (
		media            Media
		ext              = filepath.Ext(filePath)
		detectedMimeType = mime.TypeByExtension(ext)
	)

	if detectedMimeType != mimeType {
		return "", fmt.Errorf("o tipo MIME do arquivo (%s) não corresponde ao tipo fornecido: %s", detectedMimeType, mimeType)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := whatsapp.Buffer(make([]byte, 0))
	writer := multipart.NewWriter(body)

	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(filePath))},
		"Content-Type":        []string{mimeType},
	})
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	if err := writer.WriteField("messaging_product", "whatsapp"); err != nil {
		return "", err
	}

	if err := writer.WriteField("type", mimeType); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	url := whatsapp.Endpoint(whatsapp.senderID + "/media")
	res, err := whatsapp.Post(url, &goxios.RequestOpts{
		Headers: whatsapp.Headers(writer.FormDataContentType()),
		Body:    body,
	})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if err := validateStatus(res, b); err != nil {
		return "", err
	}

	if err := goxios.UnmarshalJSON(b, &media); err != nil {
		return "", err
	}

	return media.Id, nil
}

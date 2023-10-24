package notify

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadToMisskeyDrive(content MisskeyDriveUploadForm) error {
	id := uuid.New()
	apiEndpoint := fmt.Sprintf("https://%s/api/drive/files/create", content.InstanceHost)

	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	defer writer.Close()

	// ファイルフィールドを追加
	fileField, err := writer.CreateFormFile("file", id.String()+".png")
	if err != nil {
		return err
	}
	_, err = io.Copy(fileField, bytes.NewReader(content.Data))
	if err != nil {
		return err
	}

	// トークンフィールドを追加
	tokenField, err := writer.CreateFormField("i")
	if err != nil {
		return err
	}
	tokenField.Write([]byte(content.Token))

	// Content-Type ヘッダを設定
	contentType := writer.FormDataContentType()

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", apiEndpoint, requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		println(string(b))
		return fmt.Errorf("misskey API returned non-200 status code")
	}

	return nil
}

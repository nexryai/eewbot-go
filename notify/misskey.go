package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadToMisskeyDrive(content MisskeyDriveUploadForm) (MisskeyDriveFile, error) {
	var apiResp MisskeyDriveFile

	id := uuid.New()
	apiEndpoint := fmt.Sprintf("https://%s/api/drive/files/create", content.InstanceHost)

	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	defer writer.Close()

	// ファイルフィールドを追加
	fileField, err := writer.CreateFormFile("file", id.String()+".png")
	if err != nil {
		return apiResp, err
	}
	_, err = io.Copy(fileField, bytes.NewReader(content.Data))
	if err != nil {
		return apiResp, err
	}

	// トークンフィールドを追加
	tokenField, err := writer.CreateFormField("i")
	if err != nil {
		return apiResp, err
	}
	tokenField.Write([]byte(content.Token))

	// Content-Type ヘッダを設定
	contentType := writer.FormDataContentType()

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", apiEndpoint, requestBody)
	if err != nil {
		return apiResp, err
	}
	req.Header.Set("Content-Type", contentType)

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return apiResp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		println(string(b))
		return apiResp, fmt.Errorf("misskey API returned non-200 status code (/drive/files/create)")
	} else {
		fmt.Println(resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return apiResp, err
	}

	return apiResp, nil
}

func PostToMisskey(note MisskeyNote) error {
	// jsonをこねこね
	reqJson, err := json.Marshal(note)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/notes/create", note.InstanceHost), bytes.NewBuffer(reqJson))
	if err != nil {
		return err
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントを作成
	client := &http.Client{}

	// リクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// レスポンスを処理
	if resp.StatusCode == http.StatusOK {
	} else {
		b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b))
		return fmt.Errorf("misskey API returned non-200 status code (/notes/create)")
	}

	return nil
}

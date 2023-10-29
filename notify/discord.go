package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func NotifyToDiscord(content DiscordHook) error {
	postJson, _ := json.Marshal(content)

	// discord webhook_url
	hookUrl := os.Getenv("DISCORD_WEBHOOK")
	resp, err := http.Post(
		hookUrl,
		"application/json",
		bytes.NewBuffer(postJson),
	)

	println(string(postJson))

	if resp.StatusCode != 200|204 {
		b, _ := io.ReadAll(resp.Body)
		println(string(b))
		println(resp.StatusCode)
		err = fmt.Errorf("discord Webhook retrund non-200 code")
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

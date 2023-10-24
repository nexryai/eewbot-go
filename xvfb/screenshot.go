package xvfb

import (
	"bytes"
	"os/exec"
)

func TakeScreenshotOfXvfb() (*[]byte, error) {
	cmd := exec.Command("import", "-window", "root", "png:-")
	cmd.Env = []string{"DISPLAY=:99"}

	// コマンドのstdoutをキャプチャするバッファを作成
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// コマンドを実行
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	outputBytes := stdout.Bytes()

	return &outputBytes, nil
}

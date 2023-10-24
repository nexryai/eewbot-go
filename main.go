package main

import (
	"fmt"
	"github.com/nexryai/eewbot-go/notify"
	"github.com/nexryai/eewbot-go/xvfb"
	"os"
)

func main() {
	imageData, err := xvfb.TakeScreenshotOfXvfb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := notify.MisskeyDriveUploadForm{
		InstanceHost: "beta.romneko.net",
		Token:        os.Getenv("MISSKEY_TOKEN"),
		Data:         *imageData,
	}

	err = notify.UploadToMisskeyDrive(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

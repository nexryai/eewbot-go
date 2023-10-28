package main

import (
	"fmt"
	"github.com/nexryai/eewbot-go/notify"
	"github.com/nexryai/eewbot-go/quake"
	"github.com/nexryai/eewbot-go/xvfb"
	"os"
)

func main() {
	if os.Getenv("KEVI_EVENT_TYPE") == "EEW_RECEIVED" {
		imageData, err := xvfb.TakeScreenshotOfXvfb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		data := notify.MisskeyDriveUploadForm{
			InstanceHost: "miski.st",
			Token:        os.Getenv("MISSKEY_TOKEN"),
			Data:         *imageData,
		}

		driveApiResp, err := notify.UploadToMisskeyDrive(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(driveApiResp.FileID)

		intensity := os.Getenv("EEW_INTENSITY")
		var text string

		if quake.IsEmergency(intensity) {
			text = fmt.Sprintf("<center>$[x2 ⚠️緊急地震速報(EEW) 第%s報⚠️]</center>\n\n震源: **%s**\n$[fg 最大震度: 震度**%s**]",
				os.Getenv("EEW_COUNT"),
				os.Getenv("EEW_PLACE"),
				intensity)
		} else {
			text = fmt.Sprintf("<center>**⚠️緊急地震速報(EEW) 第%s報⚠️**</center>\n\n震源: **%s**\n最大震度: 震度**%s**",
				os.Getenv("EEW_COUNT"),
				os.Getenv("EEW_PLACE"),
				intensity)
		}

		note := notify.MisskeyNote{
			InstanceHost: "miski.st",
			Token:        os.Getenv("MISSKEY_TOKEN"),
			Text:         text,
			LocalOnly:    true,
			Visibility:   "public",
			FileIds:      []string{driveApiResp.FileID},
		}
		err = notify.PostToMisskey(note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

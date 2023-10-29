package main

import (
	"fmt"
	"github.com/nexryai/eewbot-go/notify"
	"github.com/nexryai/eewbot-go/quake"
	"github.com/nexryai/eewbot-go/xvfb"
	"os"
	"strconv"
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

		// 通知の準備
		reportNumInt, err := strconv.ParseInt(os.Getenv("EEW_COUNT"), 10, 16)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		q := quake.Equake{
			ReportNum:     uint16(reportNumInt),
			DispIntensity: os.Getenv("EEW_INTENSITY"),
			Place:         os.Getenv("EEW_PLACE"),
		}

		// Misskeyに通知
		var text string
		if quake.IsEmergency(q.DispIntensity) {
			text = fmt.Sprintf("<center>$[x2 ⚠️緊急地震速報(EEW) 第%d報⚠️]</center>\n\n震源: **%s**\n$[fg 最大震度: 震度**%s**]",
				q.ReportNum,
				q.Place,
				q.DispIntensity)
		} else {
			text = fmt.Sprintf("<center>**⚠️緊急地震速報(EEW) 第%d報⚠️**</center>\n\n震源: **%s**\n最大震度: 震度**%s**",
				q.ReportNum,
				q.Place,
				q.DispIntensity)
		}

		note := notify.MisskeyNote{
			InstanceHost: "miski.st",
			Token:        os.Getenv("MISSKEY_TOKEN"),
			Text:         text,
			LocalOnly:    false,
			Visibility:   "public",
			FileIds:      []string{driveApiResp.FileID},
		}

		err = notify.PostToMisskey(note)
		if err != nil {
			fmt.Println(err)
		}

		// Discordに通知
		var color int
		var discordEmbedTitle string
		var discordEmbedMessage string

		if quake.IsEmergency(q.DispIntensity) {
			color = 0xD50000

			// 本文
			discordEmbedTitle = fmt.Sprintf("最大震度%s 緊急地震速報第%d報 強い揺れに警戒",
				q.DispIntensity,
				q.ReportNum)

			discordEmbedMessage = fmt.Sprintf("強い揺れに警戒してください。%sを震源とする最大震度%sの地震が発生しました。落ち着いて行動してください。",
				q.Place,
				q.DispIntensity)

		} else {
			color = 0xffc000

			// 本文
			discordEmbedTitle = fmt.Sprintf("最大震度%s 緊急地震速報第%d報",
				q.DispIntensity,
				q.ReportNum)

			discordEmbedMessage = fmt.Sprintf("%sを震源とする最大震度%sの地震が発生しました。",
				q.Place,
				q.DispIntensity)
		}

		var discordNotify notify.DiscordHook
		discordNotify.Username = "EEW Bot"
		discordNotify.Content = "緊急地震速報を受信。"
		discordNotify.Embeds = []notify.DiscordEmbed{
			notify.DiscordEmbed{
				Title:  discordEmbedTitle,
				Desc:   discordEmbedMessage,
				Color:  color,
				Author: notify.DiscordAuthor{Name: "eewbot-go"},
				Image:  notify.DiscordImg{URL: driveApiResp.Url},
				Fields: []notify.DiscordField{
					notify.DiscordField{Name: "震源", Value: q.Place, Inline: true},
					notify.DiscordField{Name: "最大震度", Value: q.DispIntensity, Inline: true},
				},
			},
		}

		err = notify.NotifyToDiscord(discordNotify)
		if err != nil {
			os.Exit(1)
		}

	} else if os.Getenv("EEW_DEBUGMODE") == "1" {
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

		text := "Botは起動しました。"

		note := notify.MisskeyNote{
			InstanceHost: "miski.st",
			Token:        os.Getenv("MISSKEY_TOKEN"),
			Text:         text,
			LocalOnly:    true,
			Visibility:   "home",
			FileIds:      []string{driveApiResp.FileID},
		}

		err = notify.PostToMisskey(note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

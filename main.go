package main

import (
	"log"
	"os"

	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	result := `err
	timetable empty`

	date := time.Now()
	switch date.Weekday() {
	case 0:
		result = `1限：ネットワーク技術
2限：人文社会科学群
3限：情報科学`
	case 1:
		result = `1限：応用数学I
2限：制御理論I
3限：電気回路B`
	case 2:
		result = `1限：総合英語AI
2限：生涯スポーツI
3限：電子回路BI`
	case 3:
		result = `1限：電気磁気学B
2限：電気機器
3限：リベラルアーツ群
4限：情報基礎`
	case 4:
		result = `1限:数値計算法
2限：アジア文学論
3限：電子情報システム工学実験演習B
4限：電子情報誌ステう工学実験演習B`
	default:
		result = ``
	}
	message := linebot.NewTextMessage(result)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

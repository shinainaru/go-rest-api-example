package main

// run with: env PORT=8081 go run http-server.go

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
    var err error
    bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
    log.Println("Bot:", bot, " err:", err)

    port := os.Getenv("PORT")
    if port == "" {
      log.Fatal("Please specify the HTTP port as environment variable, e.g. env PORT=8081 go run http-server.go")
    }
    
    http.HandleFunc("/callback", callbackHandler)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World")
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
  events, err := bot.ParseRequest(r)
  if err != nil {
    if err == linebot.ErrInvalidSignature {
      w.WriteHeader(400)
    } else {
      w.WriteHeader(500)
    }
    return
  }

  for _, event := range events {
    if event.Type == linebot.EventTypeMessage {
    switch message := event.Message.(type) {
      // Handle only on text message
      case *linebot.TextMessage:
        // GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
        quota, err := bot.GetMessageQuota().Do()
        if err != nil {
          log.Println("Quota err:", err)
        }
        // message.ID: Msg unique ID
        // message.Text: Msg text
        if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
          log.Print(err)
        }
      // Handle only on Sticker message
      case *linebot.StickerMessage:
        var kw string
        for _, k := range message.Keywords {
          kw = kw + "," + k
        }
        outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
        if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
          log.Print(err)
        }
      }
    }
  }
}

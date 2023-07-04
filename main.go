/*
*

	此代码由Bing Ai 生成
*/
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Config struct {
	Webhook string `yaml:"webhook"`
	Secret  string `yaml:"secret"`
}

func getConfig() Config {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

type Message struct {
	MsgType  string          `json:"msgtype"`
	Text     TextMessage     `json:"text,omitempty"`
	Link     LinkMessage     `json:"link,omitempty"`
	MarkDown MarkDownMessage `json:"markdown,omitempty"`
}

type TextMessage struct {
	Content string `json:"content"`
}

type LinkMessage struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl"`
}

type MarkDownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	msgType := flag.String("type", "text", "消息类型: text, link, markdown")
	content := flag.String("content", "", "消息内容")
	title := flag.String("title", "", "消息标题")
	messageURL := flag.String("url", "", "消息链接")
	picURL := flag.String("pic", "", "图片链接")
	flag.Parse()

	message := Message{MsgType: *msgType}
	switch *msgType {
	case "text":
		message.Text = TextMessage{Content: *content}
	case "link":
		message.Link = LinkMessage{
			Title:      *title,
			Text:       *content,
			MessageURL: *messageURL,
			PicURL:     *picURL,
		}
	case "markdown":
		message.MarkDown = MarkDownMessage{
			Title: *title,
			Text:  *content,
		}
	default:
		fmt.Println("无效的消息类型")
		return
	}

	sendMessage(message)
}

func sendMessage(message Message) {
	config := getConfig()
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	signature, timestamp := getSignature(config.Secret)
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.Webhook+"&timestamp="+timestamp+"&sign="+signature, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func getSignature(secret string) (string, string) {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(stringToSign))
	data := hmac256.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(data)
	return url.QueryEscape(signature), timestamp
}

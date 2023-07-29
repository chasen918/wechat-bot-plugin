package newbing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
)

type dataQYK struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

func Bing(msg *openwechat.Message, self *openwechat.Self) {
	test, err := getMessage(msg.Content)
	if err != nil {
		msg.ReplyText("error")
		return
	}
	msg.ReplyText(test)

}

// 青云客取消息
func getMessage(msg string) (string, error) {
	qykUrl := "http://api.qingyunke.com/api.php"
	key := "free"
	appid := "0"
	qykUrl = fmt.Sprintf(qykUrl+"?key=%s&appid=%s&msg=%s", key, appid, url.QueryEscape(msg))

	client := &http.Client{}
	req, err := http.NewRequest("GET", qykUrl, nil)
	if err != nil {
		return "", err
	}
	// 自定义Header
	req.Header.Set("User-Agent", getAgent())
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "api.qingyunke.com")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(bytes))
	var dataQYK dataQYK
	if err := json.Unmarshal(bytes, &dataQYK); err != nil {
		return "", err
	}
	return msgReply(dataQYK.Content), nil
}

func getAgent() string {
	agent := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len1 := len(agent)
	return agent[r.Intn(len1)]
}

//处理消息
func msgReply(msg string) string {
	msg = strings.ReplaceAll(msg, "菲菲", "ATRI")
	msg = strings.ReplaceAll(msg, "{br}", "\n")
	return msg
}

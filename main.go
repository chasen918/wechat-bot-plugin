package main

import (
	"fmt"

	newbing "wechat-bot-plugin/plugin/newbing"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.MessageHandler = func(msg *openwechat.Message) {
		fmt.Println(msg)

		loadPlugin(msg, self)
	}

	// 获取所有的好友
	//friends, err := self.Friends()
	//fmt.Println(friends, err)

	// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

func loadPlugin(msg *openwechat.Message, self *openwechat.Self) {
	go func() {
		if OnKeyWord(msg, "青云客") {
			newbing.Bing(msg, self)
		}
	}()
}

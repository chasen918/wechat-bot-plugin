package main

import (
	"strings"

	"github.com/eatmoreapple/openwechat"
)

func OnKeyWord(msg *openwechat.Message, text string) bool {
	if msg.IsText() && strings.Contains(msg.Content, text) {
		return true
	} else {
		return false
	}
}

func OnKeyWordGroup(msg *openwechat.Message, text []string) bool {
	if !msg.IsText() {
		return false
	}
	for _, value := range text {
		if strings.Contains(msg.Content, value) {
			return true
		}
	}
	return false
}

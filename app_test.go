package main

import (
	"testing"

	"github.com/hgcncn/MiraiGo-Template/bot"
)

func TestDevice(t *testing.T) {
	b := &bot.Bot{}
	b.GenRandomDevice()
}

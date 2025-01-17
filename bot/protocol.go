package bot

import "github.com/Mrs4s/MiraiGo/client"

type protocol int

const (
	AndroidPhone = protocol(client.AndroidPhone)
	IPad         = protocol(client.IPad)
	AndroidWatch = protocol(client.AndroidWatch)
	MacOS        = protocol(client.MacOS)
)

// UseProtocol 使用协议
// 不同协议会有部分功能无法使用
// 默认为 AndroidPad
func (b *Bot) UseProtocol(p protocol) {
	device := b.Client.Device()
	device.Protocol = client.ClientProtocol(p)
}

package bot

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

// The following functions are designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnPrivateMessage(f func(qqClient *client.QQClient, event *message.PrivateMessage)) {
	bot.Client.PrivateMessageEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnPrivateMessageF(filter func(*message.PrivateMessage) bool, f func(*client.QQClient, *message.PrivateMessage)) {
	bot.OnPrivateMessage(func(client *client.QQClient, msg *message.PrivateMessage) {
		if filter(msg) {
			f(client, msg)
		}
	})
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnTempMessage(f func(qqClient *client.QQClient, event *client.TempMessageEvent)) {
	bot.Client.TempMessageEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMessage(f func(qqClient *client.QQClient, event *message.GroupMessage)) {
	bot.Client.GroupMessageEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnSelfPrivateMessage(f func(qqClient *client.QQClient, event *message.PrivateMessage)) {
	bot.Client.SelfPrivateMessageEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnSelfGroupMessage(f func(qqClient *client.QQClient, event *message.GroupMessage)) {
	bot.Client.SelfGroupMessageEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMuted(f func(qqClient *client.QQClient, event *client.GroupMuteEvent)) {
	bot.Client.GroupMuteEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMessageRecalled(f func(qqClient *client.QQClient, event *client.GroupMessageRecalledEvent)) {
	bot.Client.GroupMessageRecalledEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnFriendMessageRecalled(f func(qqClient *client.QQClient, event *client.FriendMessageRecalledEvent)) {
	bot.Client.FriendMessageRecalledEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupJoin(f func(qqClient *client.QQClient, event *client.GroupInfo)) {
	bot.Client.GroupJoinEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupLeave(f func(qqClient *client.QQClient, event *client.GroupLeaveEvent)) {
	bot.Client.GroupLeaveEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMemberJoin(f func(qqClient *client.QQClient, event *client.MemberJoinGroupEvent)) {
	bot.Client.GroupMemberJoinEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMemberLeave(f func(qqClient *client.QQClient, event *client.MemberLeaveGroupEvent)) {
	bot.Client.GroupMemberLeaveEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnMemberCardUpdated(f func(qqClient *client.QQClient, event *client.MemberCardUpdatedEvent)) {
	bot.Client.MemberCardUpdatedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupNameUpdated(f func(qqClient *client.QQClient, event *client.GroupNameUpdatedEvent)) {
	bot.Client.GroupNameUpdatedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupMemberPermissionChanged(f func(qqClient *client.QQClient, event *client.MemberPermissionChangedEvent)) {
	bot.Client.GroupMemberPermissionChangedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupInvited(f func(qqClient *client.QQClient, event *client.GroupInvitedRequest)) {
	bot.Client.GroupInvitedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnUserWantJoinGroup(f func(qqClient *client.QQClient, event *client.UserJoinGroupRequest)) {
	bot.Client.UserWantJoinGroupEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnNewFriend(f func(qqClient *client.QQClient, event *client.NewFriendEvent)) {
	bot.Client.NewFriendEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnNewFriendRequest(f func(qqClient *client.QQClient, event *client.NewFriendRequest)) {
	bot.Client.NewFriendRequestEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnDisconnected(f func(qqClient *client.QQClient, event *client.ClientDisconnectedEvent)) {
	bot.Client.DisconnectedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupNotify(f func(qqClient *client.QQClient, event client.INotifyEvent)) {
	bot.Client.GroupNotifyEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnFriendNotify(f func(qqClient *client.QQClient, event client.INotifyEvent)) {
	bot.Client.FriendNotifyEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnMemberSpecialTitleUpdated(f func(qqClient *client.QQClient, event *client.MemberSpecialTitleUpdatedEvent)) {
	bot.Client.MemberSpecialTitleUpdatedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnGroupDigest(f func(qqClient *client.QQClient, event *client.GroupDigestEvent)) {
	bot.Client.GroupDigestEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnOtherClientStatusChanged(f func(qqClient *client.QQClient, event *client.OtherClientStatusChangedEvent)) {
	bot.Client.OtherClientStatusChangedEvent.Subscribe(f)
}

// Deprecated: This function is designed to make the framework compatible with the old plug-in, and the newly developed plug-in should no longer use this method
func (bot *Bot) OnOfflineFile(f func(qqClient *client.QQClient, event *client.OfflineFileEvent)) {
	bot.Client.OfflineFileEvent.Subscribe(f)
}

package cnv

import "../msg"

// GetMessageList returns list of message for two users
func GetMessageList(user1, user2 string) []msg.Message {
	if user2 < user1 {
		user1, user2 = user2, user1
	}
	for _, c := range Conversations {
		if c.User1 == user1 && c.User2 == user2 {
			return c.Messages
		}
	}
	return make([]msg.Message, 0)
}

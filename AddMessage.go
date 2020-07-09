package cnv

import "../msg"

// AddMessage addes a new message to a two-person conversation
func AddMessage(user1, user2 string, newMsg msg.Message) {

	//sorts the two users by username
	if user2 < user1 {
		user1, user2 = user2, user1
	}

	for indx, convo := range Conversations {
		if convo.User1 == user1 && convo.User2 == user2 {
			Conversations[indx].Messages = append(Conversations[indx].Messages, newMsg)
			return
		}
	}
	//gets here if convo does not exist yet
	//create a new slice of messages
	messages := make([]msg.Message, 0)
	messages = append(messages, newMsg) //append newMsg to messages

	newConvo := Conversation{user1, user2, messages}
	Conversations = append(Conversations, newConvo)
}

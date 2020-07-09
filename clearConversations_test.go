package cnv

import (
	"testing"

	"../msg"
)

func TestClearConversation(t *testing.T) {

	//append a conversation to conversations
	user1 := "bhehar"
	user2 := "andrew"
	newMsg := msg.CreateMessage(user1, "what's up dude")
	messageSlice := make([]msg.Message, 0)
	messageSlice = append(messageSlice, newMsg)
	convo := Conversation{user1, user2, messageSlice}
	Conversations = append(Conversations, convo)

	clearConversations()

	got := len(Conversations)
	expected := 0

	if got != expected {
		t.Errorf("got %v expected %v", got, expected)
	}
}

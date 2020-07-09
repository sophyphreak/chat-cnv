package cnv

import (
	"testing"

	"../msg"
)

func TestGetMessages(t *testing.T) {
	defer clearConversations()

	var m msg.Message
	m.Username = "bali"
	m.Body = "wzzzp"

	var c Conversation
	c.User1 = "bali"
	c.User2 = "jasmine"
	c.Messages = append(c.Messages, m)

	Conversations = append(Conversations, c)

	var messages []msg.Message
	messages = GetMessageList(c.User2, c.User1)
	if len(messages) != 1 {
		t.Errorf("Expected length of messages to be 1 but instead received %d", len(messages))
	}
	if messages[0].Username != c.User1 {
		t.Errorf("Expected %s as messsage username but instead received %s", c.User1, messages[0].Username)
	}
	if messages[0].Body != m.Body {
		t.Errorf("Expected %s as message body but instead received %s", m.Body, messages[0].Body)
	}
}

package pubsub

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const (
	PUBLISH = "publish"
)

func autoId() string {
	return uuid.NewV4().String()
}

type PubSub struct {
	Clients []Client
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}

type Message struct {
	Action    string          `json:"action"`
	MeetingID string          `json:"meetingId"`
	Message   json.RawMessage `json:"message"`
}

func (ps *PubSub) AddClient(client Client) *PubSub {

	ps.Clients = append(ps.Clients, client)

	//fmt.Println("adding new client to the list", client.Id, len(ps.Clients))

	payload := []byte(client.Id)

	client.Connection.WriteMessage(1, payload)
	return ps
}

func (ps *PubSub) RemoveClient(client Client) *PubSub {

	for index, c := range ps.Clients {

		if c.Id == client.Id {
			ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
		}

	}

	return ps
}

func (ps *PubSub) Publish(id string, message []byte, excludeClient *Client) {
	for _, sub := range ps.Clients {
		if excludeClient != nil && excludeClient.Id == sub.Id {

		} else {
			//sub.Client.Connection.WriteMessage(1, message)
			sub.Send(message)
		}
	}

}

func (client *Client) Send(message []byte) error {

	return client.Connection.WriteMessage(1, message)

}

func (ps *PubSub) HandleReceiveMessage(client Client, messageType int, payload []byte) *PubSub {

	m := Message{}

	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Println("This is not correct message payload")
		return ps
	}

	switch m.Action {

	case PUBLISH:
		ps.Publish(m.MeetingID, m.Message, &client)
	default:
		break
	}

	return ps
}

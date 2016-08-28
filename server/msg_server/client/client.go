package client

import (
	"github.com/oikomi/FishChatServer2/libnet"
	"github.com/oikomi/FishChatServer2/protocol"
)

type Client struct {
	Session *libnet.Session
}

func New(session *libnet.Session) (c *Client) {
	c = &Client{
		Session: session,
	}
	return
}

func (c *Client) Parse(cmd uint32, reqData []byte) {

}
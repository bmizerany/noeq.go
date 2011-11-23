package noeq

import (
	"encoding/binary"
	"errors"
	"math/rand"
	"net"
	"sync"
)

var ErrNoAddrs = errors.New("noeq: no addresses provided")

type Client struct {
	mu    sync.Mutex
	cn    net.Conn
	addrs []string
}

func New(addrs ...string) (*Client, error) {
	if len(addrs) == 0 {
		return nil, ErrNoAddrs
	}
	return &Client{addrs: addrs}, nil
}

func (c *Client) connect() (err error) {
	n := rand.Intn(len(c.addrs) - 1)
	c.cn, err = net.Dial("tcp", c.addrs[n])
	return
}

func (c *Client) Gen(n uint8) (ids []uint64, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	defer func() {
		if err != nil {
			c.cn = nil
		}
	}()

	if c.cn == nil {
		err = c.connect()
		if err != nil {
			return
		}
	}

	ids = make([]uint64, n)

	_, err = c.cn.Write([]byte{n})
	if err != nil {
		return
	}

	err = binary.Read(c.cn, binary.BigEndian, &ids)
	return
}

func (c *Client) GenOne() (uint64, error) {
	ids, err := c.Gen(1)
	if len(ids) == 0 {
		return 0, err
	}
	return ids[0], nil
}

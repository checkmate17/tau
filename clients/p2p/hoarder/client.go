package p2p

import (
	"context"
	"errors"
	"fmt"
	"time"

	moodyBlues "bitbucket.org/taubyte/go-moody-blues"
	"github.com/taubyte/go-interfaces/moody"
	iface "github.com/taubyte/go-interfaces/services/hoarder"
	protocolCommon "github.com/taubyte/odo/protocols/common"
	"github.com/taubyte/p2p/peer"
	client "github.com/taubyte/p2p/streams/client"
)

var _ iface.Client = &Client{}

var (
	MinPeers                 = 0
	MaxPeers                 = 2
	DefaultGeoBeaconInterval = 5 * time.Minute
	ErrorGeoBeaconStopped    = errors.New("geoBeacon Stopped")
	logger                   moody.Logger
)

type Client struct {
	client *client.Client
}

func init() {
	logger, _ = moodyBlues.New("hoarder.p2p.client")
}

func New(ctx context.Context, node peer.Node) (*Client, error) {
	var (
		c   Client
		err error
	)

	if c.client, err = client.New(ctx, node, nil, protocolCommon.HoarderProtocol, MinPeers, MaxPeers); err != nil {
		logger.Errorf(fmt.Sprintf("API client creation failed: %v", err))
		return nil, err
	}
	return &c, nil
}

func (c *Client) Close() {
	c.client.Close()
}

/* Peer */

type Peer struct {
	Client
	Id string
}

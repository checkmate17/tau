package mocks

import (
	"github.com/taubyte/go-interfaces/services/substrate"
	"github.com/taubyte/go-interfaces/services/tns"
	httpSrv "github.com/taubyte/http"
	"github.com/taubyte/p2p/peer"
)

func Node(node peer.Node) option {
	return func(mn *mockedSubstrate) error {
		mn.node = node
		return nil
	}
}

func TNS(client tns.Client) option {
	return func(mn *mockedSubstrate) error {
		mn.tns = client
		return nil
	}
}

func HTTP(service httpSrv.Service) option {
	return func(mn *mockedSubstrate) error {
		mn.http = service
		return nil
	}
}

func SmartOps(service substrate.SmartOpsService) option {
	return func(mn *mockedSubstrate) error {
		mn.smartOps = service
		return nil
	}
}

func Branch(branch string) option {
	return func(ms *mockedSubstrate) error {
		ms.branch = branch
		return nil
	}
}

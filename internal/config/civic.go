package config

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type Civic struct {
	GatekeeperNetworkId  *big.Int
	GatewayTokenContract common.Address
	EthereumRpc          *ethclient.Client
	XDCRpc               *ethclient.Client
	PolygonRpc           *ethclient.Client
	ArbitrumRpc          *ethclient.Client
	SkipSigCheck         bool
}

type civic struct {
	// GatekeeperNetworkId ID of the gatekeeperNetwork on gateway contract. You can find info about it in Civic doc.
	// aka gatekeeper_network_slot_id / network (uint256) on chain
	GatekeeperNetworkId *big.Int `fig:"gatekeeper_network_id"`
	// GatewayTokenContract 0xF65b6396dF6B7e2D8a6270E3AB6c7BB08BAEF22E
	GatewayTokenContract string `fig:"gateway_token_contract,required"`
	EthRpcURL            string `fig:"eth_rpc_url,required"`
	XDCRpcURL            string `fig:"xdc_rpc_url,required"`
	PolygonRpcURL        string `fig:"polygon_rpc_url,required"`
	ArbitrumRpcURL       string `fig:"arbitrum_rpc_url,required"`
	SkipSigCheck         bool   `fig:"skip_sig_check,required"`
}

func (c *config) Civic() *Civic {
	return c.civic.Do(func() interface{} {
		cfg := civic{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "civic")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return parseCivicConfig(&cfg)
	}).(*Civic)
}

func parseCivicConfig(cfg *civic) *Civic {
	ethClient, err := ethclient.Dial(cfg.EthRpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create an Ethereum client"))
	}

	polygonClient, err := ethclient.Dial(cfg.PolygonRpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create a Polygon client"))
	}

	arbitrumClient, err := ethclient.Dial(cfg.ArbitrumRpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create an Arbitrum client"))
	}

	xdcClient, err := ethclient.Dial(cfg.XDCRpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create a XDC client"))
	}

	if !common.IsHexAddress(cfg.GatewayTokenContract) {
		panic(errors.New("failed to parse gateway token contract address"))
	}

	return &Civic{
		GatekeeperNetworkId:  cfg.GatekeeperNetworkId,
		GatewayTokenContract: common.HexToAddress(cfg.GatewayTokenContract),
		EthereumRpc:          ethClient,
		XDCRpc:               xdcClient,
		PolygonRpc:           polygonClient,
		ArbitrumRpc:          arbitrumClient,
		SkipSigCheck:         cfg.SkipSigCheck,
	}
}

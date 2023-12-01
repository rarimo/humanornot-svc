package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Kleros struct {
	EthereumRpc             *ethclient.Client
	ProofOfHumanityContract common.Address
}

type kleros struct {
	EthRpcURL               string `fig:"eth_rpc_url,required"`
	ProofOfHumanityContract string `fig:"proof_of_humanity_contract,required"`
}

func (c *config) Kleros() *Kleros {
	return c.civic.Do(func() interface{} {
		cfg := kleros{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "kleros")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return parseKlerosConfig(&cfg)
	}).(*Kleros)
}

func parseKlerosConfig(cfg *kleros) *Kleros {
	ethClient, err := ethclient.Dial(cfg.EthRpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create an Ethereum client"))
	}

	if !common.IsHexAddress(cfg.ProofOfHumanityContract) {
		panic(errors.New("failed to parse gateway token contract address"))
	}

	return &Kleros{
		EthereumRpc:             ethClient,
		ProofOfHumanityContract: common.HexToAddress(cfg.ProofOfHumanityContract),
	}
}

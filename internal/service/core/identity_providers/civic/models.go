package civic

type ChainName string

func (c ChainName) String() string {
	return string(c)
}

const (
	EthereumChainName ChainName = "ethereum"
	PolygonChainName  ChainName = "polygon"
	ArbitrumChainName ChainName = "arbitrum"
	XDCChainName      ChainName = "xdc"
)

var chainNameFromString = map[string]ChainName{
	"ethereum": EthereumChainName,
	"polygon":  PolygonChainName,
	"arbitrum": ArbitrumChainName,
	"xdc":      XDCChainName,
}

type VerificationData struct {
	ChainName ChainName `json:"chain_name"`
	Signature string    `json:"signature"`
	Address   string    `json:"address"`
}

// Client is an interface for the web3 RPC API.
type Client interface {
	// GetBalance returns the balance for an address at the given block number (nil for latest).
	GetBalance(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error)
// GetCode returns the code for an address at the given block number (nil for latest).
	GetCode(ctx context.Context, address string, blockNumber *big.Int) ([]byte, error)
GetBlockByNumber(ctx context.Context, number *big.Int, includeTxs bool) (*Block, error)
	// GetBlockByHash returns block details for the given hash, optionally include full transaction details.
GetBlockByHash(ctx context.Context, hash string, includeTxs bool) (*Block, error)
	// GetTransactionByHash returns transaction details for a hash.
GetTransactionByHash(ctx context.Context, hash common.Hash) (*Transaction, error)
	// GetSnapshot returns the latest clique snapshot.
GetSnapshot(ctx context.Context) (*Snapshot, error)
	// GetID returns unique identifying information for the network.
GetChainID(ctx context.Context) (*big.Int, error)
	// GetNetworkID returns the network id.
GetNetworkID(ctx context.Context) (*big.Int, error)
	// GetGasPrice returns a suggested gas price.
GetGasPrice(ctx context.Context) (*big.Int, error)
	// GetPendingTransactionCount returns the transaction count including pending txs.
// This value is also the next legal nonce.
	GetPendingTransactionCount(ctx context.Context, account common.Address) (uint64, error)
// SendRawTransaction sends the signed raw transaction bytes.
	SendRawTransaction(ctx context.Context, tx []byte) error
// Call executes a call without submitting a transaction.
	Call(ctx context.Context, msg CallMsg) ([]byte, error)
Close()
	SetChainID(*big.Int)
}

// Dial returns a new client backed by dialing url (supported schemes "http", "https", "ws" and "wss").
func Dial(url string) (Client, error) {
r, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return NewClient(r), nil
}

// NewClient returns a new client backed by an existing rpc.Client.
func NewClient(r *rpc.Client) Client {
	return &client{r: r}
}

type client struct {
r       *rpc.Client
	chainID atomic.Value
}

func (c *client) Close() {
c.r.Close()
}

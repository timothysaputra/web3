// Client is an interface for the web3 RPC API.
type Client interface {
	// GetBalance returns the balance for an address at the given block number (nil for latest).
	GetBalance(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error)
// GetCode returns the code for an address at the given block number (nil for latest).
	GetCode(ctx context.Context, address string, blockNumber *big.Int) ([]byte, error)
GetBlockByNumber(ctx context.Context, number *big.Int, includeTxs bool) (*Block, error)
	// GetBlockByHash returns block details for the given hash, optionally include full transaction details.

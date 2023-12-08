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


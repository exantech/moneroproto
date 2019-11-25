package moneroproto

import "github.com/exantech/moneroutil"

type GetHashesFastRequest struct {
	Client      []byte `monerobinkv:"client"`
	BlockIds    []byte `monerobinkv:"block_ids"`
	StartHeight uint64 `monerobinkv:"start_height"`
}

func (g *GetHashesFastRequest) SetHashes(hashes []moneroutil.Hash) {
	g.BlockIds = HashesToByteSlice(hashes)
}

func (g *GetHashesFastRequest) GetHashes() (error, []moneroutil.Hash) {
	return ByteSliceToHashes(g.BlockIds)
}

type GetHashesFastResponse struct {
	BlockIds      []byte `monerobinkv:"m_block_ids"`
	StartHeight   uint64 `monerobinkv:"start_height"`
	CurrentHeight uint64 `monerobinkv:"current_height"`
	Status        []byte `monerobinkv:"status"`
	Untrusted     bool   `monerobinkv:"untrusted"`
	Credits       uint64 `monerobinkv:"credits"`
	TopHash       []byte `monerobinkv:"top_hash"`
}

func (g *GetHashesFastResponse) SetHashes(hashes []moneroutil.Hash) {
	g.BlockIds = HashesToByteSlice(hashes)
}

func (g *GetHashesFastResponse) GetHashes() (error, []moneroutil.Hash) {
	return ByteSliceToHashes(g.BlockIds)
}

type GetBlocksFastRequest struct {
	Client      []byte `monerobinkv:"client"`
	BlockIds    []byte `monerobinkv:"block_ids"`
	StartHeight uint64 `monerobinkv:"start_height"`
	Prune       bool   `monerobinkv:"prune"`
	NoMinerTx   bool   `monerobinkv:"no_miner_tx"`
}

func (g *GetBlocksFastRequest) SetHashes(hashes []moneroutil.Hash) {
	g.BlockIds = HashesToByteSlice(hashes)
}

func (g *GetBlocksFastRequest) GetHashes() (error, []moneroutil.Hash) {
	return ByteSliceToHashes(g.BlockIds)
}

type BlockCompleteEntry struct {
	Pruned      bool     `monerobinkv:"pruned"`
	Block       []byte   `monerobinkv:"block"`
	BlockWeight uint64   `monerobinkv:"block_weight"`
	Txs         [][]byte `monerobinkv:"txs"`
}

type TxOutputIndices struct {
	Indices []uint64 `monerobinkv:"indices"`
}

type BlockOutputIndices struct {
	Indices []TxOutputIndices `monerobinkv:"indices"`
}

type GetBlocksFastResponse struct {
	Blocks        []BlockCompleteEntry `monerobinkv:"blocks"`
	StartHeight   uint64               `monerobinkv:"start_height"`
	CurrentHeight uint64               `monerobinkv:"current_height"`
	Status        []byte               `monerobinkv:"status"`
	OutputIndices []BlockOutputIndices `monerobinkv:"output_indices"`
	Untrusted     bool                 `monerobinkv:"untrusted"`
	Credits       uint64               `monerobinkv:"credits"`
	TopHash       []byte               `monerobinkv:"top_hash"`
}

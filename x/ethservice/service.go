package ethservice

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

type EthService struct {
	client *rpc.Client
}

func NewEthService(rpcURL string) (*EthService, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &EthService{client: client}, nil
}

// Now you can add methods for various Ethereum RPC interactions you need

func (es *EthService) GetLatestBlockNumber() (uint64, error) {
	var result string
	err := es.client.Call(&result, "eth_blockNumber")
	if err != nil {
		return 0, err
	}
	// Convert result (in hex) to uint64
	blockNumber, err := strconv.ParseUint(result[2:], 16, 64)
	return blockNumber, err
}

func (es *EthService) GetBlockHashByNumber(blockNumber uint64) (string, error) {
	var block map[string]interface{}

	// "latest" can be replaced with a specific block number in hex format
	err := es.client.Call(&block, "eth_getBlockByNumber", toHex(blockNumber), true)
	if err != nil {
		return "", err
	}

	// Extracting the hash from the block's details
	if hash, ok := block["hash"].(string); ok {
		return hash, nil
	}

	return "", errors.New("block hash not found")
}

// Helper function to convert uint64 to Ethereum's hex format
func toHex(number uint64) string {
	return fmt.Sprintf("0x%x", number)
}

func (es *EthService) GetStorageAt(address string, position string, blockTag string) (string, error) {
	var result string

	err := es.client.Call(&result, "eth_getStorageAt", address, position, blockTag)
	if err != nil {
		return "", err
	}
	return result, nil
}

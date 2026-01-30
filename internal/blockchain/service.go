package blockchain

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"log"
	"math/big"
	"os"

	"muleshield/internal/blockchain/contract" // Ye wo file hai jo abhi generate hui

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct {
	client     *ethclient.Client
	contract   *contract.MuleRegistry
	privateKey *ecdsa.PrivateKey
	auth       *bind.TransactOpts
}

func NewService() (*Service, error) {
	// 1. Connect to Sepolia RPC
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_RPC_URL"))
	if err != nil {
		return nil, err
	}

	// 2. Load Private Key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	// 3. Load Contract Address
	address := common.HexToAddress(os.Getenv("MULE_CONTRACT_ADDRESS"))
	instance, err := contract.NewMuleRegistry(address, client)
	if err != nil {
		return nil, err
	}

	// 4. Create Auth (Transaction Signer)
	chainID, _ := client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	return &Service{
		client:     client,
		contract:   instance,
		privateKey: privateKey,
		auth:       auth,
	}, nil
}

func (s *Service) ReportMule(txID, sender, receiver, amountStr, location, reason string) error {
	// Accounts ko hash karein (kyunki Smart Contract bytes32 leta hai)
	senderHash := sha256.Sum256([]byte(sender))
	receiverHash := sha256.Sum256([]byte(receiver))

	// Amount ko BigInt mein convert karein
	amount := new(big.Int)
	amount.SetString(amountStr, 10)

	// Transaction Bhejein
	tx, err := s.contract.ReportMule(s.auth, txID, senderHash, receiverHash, amount, location, reason)
	if err != nil {
		return err
	}

	log.Printf(" [BLOCKCHAIN] Transaction Submitted! Hash: %s", tx.Hash().Hex())
	return nil
}
// Add this to internal/blockchain/service.go

func (s *Service) IsWalletBlacklisted(accountID string) (bool, error) {
	// 1. Hash the account ID (Same logic as ReportMule)
	hash := sha256.Sum256([]byte(accountID))

	// 2. Convert to [32]byte array for Solidity
	var hash32 [32]byte
	copy(hash32[:], hash[:])

	// 3. Call the Smart Contract (Read-Only)
	// 'nil' is passed for TransactOpts because we are only reading, not spending gas.
	isMule, err := s.contract.IsAccountFlagged(nil, hash32)
	if err != nil {
		return false, err
	}

	return isMule, nil
}
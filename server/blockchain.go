package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	tracking "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/config/Tracking"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func blockchainMain(PluginName string, WorkflowId string) {

	// fmt.Println("qualified:", results["qualified"])
	fmt.Println("name:", PluginName)
	fmt.Println("workflow:", WorkflowId)
	fmt.Println("Connecting to Ethereum blockchain...")
	// Connect to the Hardhat blockchain (Default: port 8545)
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_holesky")
	if err != nil {
		log.Fatal("Error connecting to blockchain:", err)
	}
	fmt.Println("Connected to Ethereum blockchain")

	// Replace with your actual deployed contract address
	contractAddress := common.HexToAddress("0x91719E51c3100D792DdcF4af35B661Fb60754E87") // Corrected

	// Load contract instance
	instance, err := tracking.NewTracking(contractAddress, client)
	if err != nil {
		log.Fatal("Error loading contract:", err)
	}

	fmt.Println("Smart contract loaded successfully!")

	// Convert qualified count to *big.Int
	// var qualifiedStr string
	// if val, exists := results["qualified"]; exists {
	// 	qualifiedStr = val
	// } else if val, exists := results["totalCount"]; exists {
	// 	qualifiedStr = val
	// } else {
	// 	log.Fatal("Neither 'qualified' nor 'totalCount' found in results map")
	// }
	// qualifiedCount, ok := new(big.Int).SetString(qualifiedStr, 10)
	// if !ok {
	// 	log.Fatal("Error converting qualified count to big.Int:", qualifiedStr)
	// }
	// fmt.Println("Qualified count:", qualifiedCount)

	// Call functions
	createShipment(client, instance, PluginName, WorkflowId)
	//getAllShipments(instance)
}

func createShipment(client *ethclient.Client, instance *tracking.Tracking, PluginName string, WorkflowId string) {
	// Replace with the private key of the sender
	privateKey, err := crypto.HexToECDSA("12022630c9d2eb7d4335a831f6268d78f9b4192e978e54d57d0e0401eff8b165") // Replace with a real Hardhat test account private key
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// Derive sender's public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	senderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Sender Address:", senderAddress.Hex())

	// Create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(17000)) // Use Hardhat's chain ID
	if err != nil {
		log.Fatal("Failed to create auth:", err)
	}

	// Manually set gas price to avoid `eth_maxPriorityFeePerGas` error
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get gas price:", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000) // Adjust gas limit as needed

	// Define shipment details
	shipmentId := WorkflowId // Use WorkflowId as shipment ID
	receiver := common.HexToAddress("0x9ef57661C968aCe446EB1B0BA1A3fBf607AEC12A")
	completedStep := PluginName
	acceptedCount := big.NewInt(100) // *big.Int

	// 🔁 Call smart contract function
	tx, err := instance.CreateShipment(auth, shipmentId, receiver, completedStep, acceptedCount)
	if err != nil {
		log.Fatal("Transaction failed:", err)
	}

	fmt.Println("Shipment created successfully!")
	fmt.Println("Transaction hash:", tx.Hash().Hex())
}

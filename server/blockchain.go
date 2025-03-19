package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	tracking "Coconut-Peat-Supply-chain_core_system/config/Tracking"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func blockchainMain() {
	// Connect to the Hardhat blockchain (Default: port 8545)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Error connecting to blockchain:", err)
	}
	fmt.Println("Connected to Ethereum blockchain")

	// Replace with your actual deployed contract address
	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3") // Corrected

	// Load contract instance
	instance, err := tracking.NewTracking(contractAddress, client)
	if err != nil {
		log.Fatal("Error loading contract:", err)
	}

	fmt.Println("Smart contract loaded successfully!")

	// Call functions
	createShipment(client, instance)
	getAllShipments(instance)
}

func createShipment(client *ethclient.Client, instance *tracking.Tracking) {
	// Replace with the private key of the sender
	privateKey, err := crypto.HexToECDSA("df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e") // Replace with a real Hardhat test account private key
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// Derive sender's public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	senderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Sender Address:", senderAddress.Hex())

	// Create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // Use Hardhat's chain ID
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
	receiver := common.HexToAddress("0xdD2FD4581271e230360230F9337D5c0430Bf44C0") // Replace with actual Hardhat test account
	pickupTime := big.NewInt(1678900000)                                          // Example timestamp
	distance := big.NewInt(100)
	price := big.NewInt(6000000000000000000) // 0.05 ETH (in wei)

	// Send ETH along with the transaction
	auth.Value = price // Must match the price in the contract

	// Send transaction to create shipment
	tx, err := instance.CreateShipment(auth, receiver, pickupTime, distance, price)
	if err != nil {
		log.Fatal("Transaction failed:", err)
	}

	fmt.Println("Shipment created! TX Hash:", tx.Hash().Hex())
}

func getAllShipments(instance *tracking.Tracking) {
	shipments, err := instance.GetAllTransactions(nil)
	if err != nil {
		log.Fatal("Failed to retrieve shipments:", err)
	}

	fmt.Println("List of Shipments:")
	for i, s := range shipments {
		fmt.Printf("[%d] Sender: %s, Receiver: %s, Distance: %d km, Price: %d wei, Status: %d\n",
			i, s.Sender.Hex(), s.Receiver.Hex(), s.Distance, s.Price, s.Status)
	}
}

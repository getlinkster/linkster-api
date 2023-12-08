package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	qrcode "github.com/skip2/go-qrcode"
)

func createQrCode(url string) ([]byte, error) {
	qrCodeString, err := sendGetQrRequest(url)
	if err != nil {
		return nil, err
	}
	fmt.Println("QR code string: ", qrCodeString)

	png, err := qrcode.Encode(qrCodeString, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func CanMakeNewConnection(signature string) (bool, error) {
	message := "TEST_SIGNED_MESSAGE"
	address, err := VerifyMessage(context.Background(), message, signature)
	if err != nil {
		return false, err
	}

	client, err := ethclient.Dial(config.MumbaiRpcUrl)
	if err != nil {
		return false, err
	}

	subscriptionRecorder, err := NewSubscriptionRecorder(common.HexToAddress(config.SubscriptionRecorderMumbaiAddress), client)
	if err != nil {
		return false, err
	}

	return subscriptionRecorder.HasValidSubscription(nil, address, 1)
}

func VerifyMessage(ctx context.Context, message string, signedMessage string) (common.Address, error) {
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(signedMessage)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA), nil
}

func HasValidSubscription(signature string) (bool, error) {
	message := "TEST_SIGNED_MESSAGE"
	address, err := VerifyMessage(context.Background(), message, signature)
	if err != nil {
		return false, err
	}

	client, err := ethclient.Dial(config.MumbaiRpcUrl)
	if err != nil {
		return false, err
	}
	subscriptionRecorder, err := NewSubscriptionRecorder(common.HexToAddress(config.SubscriptionRecorderMumbaiAddress), client)
	if err != nil {
		return false, err
	}

	return subscriptionRecorder.HasValidSubscription(nil, address, 0)
}

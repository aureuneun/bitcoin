package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/aureuneun/bitcoin/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

const (
	fileName string = "bitcoin.wallet"
)

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() *ecdsa.PrivateKey {
	bytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err := x509.ParseECPrivateKey(bytes)
	utils.HandleErr(err)
	return key
}

func encodeBigInts(x, y []byte) string {
	z := append(x, y...)
	return fmt.Sprintf("%x", z)
}

func addressFromKey(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func Sign(payload string, w *wallet) string {
	bytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, bytes)
	utils.HandleErr(err)
	return encodeBigInts(r.Bytes(), s.Bytes())
}

func decodeBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	fBytes := bytes[:len(bytes)/2]
	eBytes := bytes[len(bytes)/2:]
	bigF, bigE := big.Int{}, big.Int{}
	bigF.SetBytes(fBytes)
	bigE.SetBytes(eBytes)
	return &bigF, &bigE, nil
}

func Verity(payload, address, signature string) bool {
	r, s, err := decodeBigInts(signature)
	utils.HandleErr(err)
	x, y, err := decodeBigInts(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	bytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	return ecdsa.Verify(&publicKey, bytes, r, s)
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = addressFromKey(w.privateKey)
	}
	return w
}

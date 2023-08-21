package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/bnb-chain/go-sdk/common/types"
	"github.com/bnb-chain/go-sdk/keys"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/go-chain/go-tron/account"
)

type Account struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

type Response struct {
	Success bool  `json:"success"`
	Data    any   `json:"data"`
	Error   error `json:"error"`
}

func newErrorResponse(err error) *Response {
	return &Response{
		Success: false,
		Error:   err,
	}
}

func newSuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

func GenerateETHAccount() (*Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	account := new(Account)
	account.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	account.PrivateKey = fmt.Sprintf("%x", key.D.Bytes())

	return account, nil
}

func GenerateTronAccount() *Account {
	tronAcc := account.NewLocalAccount()
	account := new(Account)
	account.Address = tronAcc.Address().ToBase16()
	account.PrivateKey = tronAcc.PrivateKey()
	return account
}

func GenerateBSCAccount() (*Account, error) {
	privKey, err := generateRandomString(64)
	if err != nil {
		return nil, err
	}
	keyManager, err := keys.NewPrivateKeyManager(privKey)
	if err != nil {
		return nil, err
	}
	addr := keyManager.GetAddr().String()
	return &Account{
		PrivateKey: privKey,
		Address:    addr,
	}, nil
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789abcdef"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func main() {
	types.SetNetwork(types.TestNetwork)

	r := gin.Default()
	accountRouter := r.Group("/generate/account")
	accountRouter.GET("/eth", func(c *gin.Context) {
		ethAccount, err := GenerateETHAccount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, newErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, newSuccessResponse(ethAccount))
	})

	accountRouter.GET("/tron", func(c *gin.Context) {
		tronAccount := GenerateTronAccount()
		c.JSON(http.StatusOK, newSuccessResponse(tronAccount))
	})

	accountRouter.GET("/bsc", func(c *gin.Context) {
		bscAccount, err := GenerateBSCAccount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, newErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, newSuccessResponse(bscAccount))
	})

	r.Run(":8080")
}

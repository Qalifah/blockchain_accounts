package main

import "testing"

func TestGenerateETHAccount(t *testing.T) {
	ethAccount, err := GenerateETHAccount()
	if err != nil {
		t.Error("Unable to generate ETH Account", err)
	}
	t.Logf("ETH Address: %s", ethAccount.Address)
	t.Logf("ETH PrivateKey: %s", ethAccount.PrivateKey)
}

func TestGenerateTronAccount(t *testing.T) {
	tronAccount := GenerateTronAccount()
	t.Logf("Tron Address: %s", tronAccount.Address)
	t.Logf("Tron PrivateKey: %s", tronAccount.PrivateKey)
}

func TestGenerateBSCAccount(t *testing.T) {
	bscAccount, err := GenerateBSCAccount()
	if err != nil {
		t.Error("Unable to generate BSC Account", err)
	}
	t.Logf("BSC Address: %s", bscAccount.Address)
	t.Logf("BSC PrivateKey: %s", bscAccount.PrivateKey)
}
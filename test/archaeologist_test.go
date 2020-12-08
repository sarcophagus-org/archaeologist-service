package test

import (
	"context"
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type ArchTestSuite struct {
	suite.Suite
	auth               *bind.TransactOpts
	sarcoAddress       common.Address
	sarcoTokenAddress  common.Address
	gAlloc             core.GenesisAlloc
	sim                *backends.SimulatedBackend
	arch               *models.Archaeologist
	sarcophagus        *Sarcophagus
	sarcoToken         *SarcophagusToken
	sarcophagusSession SarcophagusSession
}

func TestRunArchTestSuite(t *testing.T) {
	suite.Run(t, new(ArchTestSuite))
}

func (s *ArchTestSuite) DeploySarcoToken() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	s.sarcoTokenAddress = auth.From
	alloc := map[common.Address]core.GenesisAccount{
		s.sarcoTokenAddress: {Balance: big.NewInt(10000000000)},
	}

	blockGasLimit := uint64(18446744073709551615)
	sim := backends.NewSimulatedBackend(alloc, blockGasLimit)
	sarcoTokenAddress, _, sarcoToken, err := DeploySarcophagusToken(auth, sim, big.NewInt(10000000000), "SarcoToken", "SarcoToken")
	if err != nil {
		s.T().Logf("Error deploying sarco token: %v", err)
	}
	s.sarcoTokenAddress = sarcoTokenAddress
	s.sarcoToken = sarcoToken
	s.Nil(err)
	s.sim.Commit()
}

func (s *ArchTestSuite) SetupTest() {
	key, _ := crypto.GenerateKey()
	s.auth = bind.NewKeyedTransactor(key)

	s.sarcoAddress = s.auth.From

	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.sarcoAddress: {Balance: big.NewInt(10000000000)},
	}

	blockGasLimit := uint64(18446744073709551615)
	s.sim = backends.NewSimulatedBackend(s.gAlloc, blockGasLimit)

	_, _, sarcophagus, e := DeploySarcophagus(s.auth, s.sim, s.sarcoTokenAddress)
	s.sarcophagus = sarcophagus
	s.Nil(e)
	s.sim.Commit()
	s.sarcophagusSession = SarcophagusSession{
		Contract:     s.sarcophagus,
		TransactOpts: *s.auth,
		CallOpts: bind.CallOpts{
			From:    s.auth.From,
			Context: context.Background(),
		},
	}
	s.InitArchaeologist(key)
}

func (s *ArchTestSuite) InitArchaeologist(privKey *ecdsa.PrivateKey) {
	wallet, _ := hdwallet.NewFromMnemonic("index cupboard city neither axis spot thumb pet rabbit stuff culture project top fault wisdom")
	var sarcoSession = contracts.SarcophagusSession{}
	accountIndex := 0
	currentPrivateKey := hdw.PrivateKeyFromIndex(wallet, accountIndex)
	currentPublicKeyBytes := hdw.PublicKeyBytesFromIndex(wallet, accountIndex)

	s.arch = &models.Archaeologist{
		Client:                nil,
		ArweaveWallet:         nil,
		ArweaveTransactor:     nil,
		PrivateKey:            privKey,
		CurrentPublicKeyBytes: currentPublicKeyBytes,
		CurrentPrivateKey:     currentPrivateKey,
		ArchAddress:           utility.PrivateKeyToAddress(privKey),
		PaymentAddress:        utility.PrivateKeyToAddress(privKey),
		SarcoAddress:          s.sarcoAddress,
		SarcoSession:          sarcoSession,
		SarcoTokenAddress:     s.sarcoTokenAddress,
		TokenSession:          contracts.TokenSession{},
		FreeBond:              big.NewInt(1000),
		FeePerByte:            big.NewInt(1),
		MinBounty:             big.NewInt(20),
		MinDiggingFee:         big.NewInt(5),
		MaxResurectionTime:    big.NewInt(1641859200),
		Endpoint:              "192.168.1.1",
		FilePort:              "7000",
		Mnemonic:              "index cupboard city neither axis spot thumb pet rabbit stuff culture project top fault wisdom",
		Wallet:                wallet,
		AccountIndex:          accountIndex,
		Server:                nil,
		Sarcophaguses:         nil,
		FileHandlers:          nil,
	}
}

func (s *ArchTestSuite) TestArchaeologistCount() {
	count, err := s.sarcophagusSession.ArchaeologistCount()
	s.Equal(int64(0), count.Int64())
	s.Nil(err)
}

func (s *ArchTestSuite) TestArchaeologistRegister() {
	_, err := s.sarcophagusSession.RegisterArchaeologist(
		s.arch.CurrentPublicKeyBytes,
		s.arch.Endpoint,
		s.arch.PaymentAddress,
		s.arch.FeePerByte,
		s.arch.MinBounty,
		s.arch.MinDiggingFee,
		s.arch.MaxResurectionTime,
		s.arch.FreeBond,
	)
	s.Nil(err)

	count, err := s.sarcophagusSession.ArchaeologistCount()
	s.Equal(int64(1), count.Int64())
	s.Nil(err)
}

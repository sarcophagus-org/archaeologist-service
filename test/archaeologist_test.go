package test

import (
	"context"
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
	auth              *bind.TransactOpts
	sarcoAddress      common.Address
	sarcoTokenAddress common.Address
	gAlloc            core.GenesisAlloc
	sim               *backends.SimulatedBackend
	sarcophagus       *Sarcophagus
	sarcoToken        *SarcophagusToken
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
	s.sarcophagusSession = SarcophagusSession {
		Contract: s.sarcophagus,
		TransactOpts: *s.auth,
		CallOpts: bind.CallOpts{
			From:    s.auth.From,
			Context: context.Background(),
		},
	}
}

func (s *ArchTestSuite) TestArchaeologistCount() {
	count, err := s.sarcophagusSession.ArchaeologistCount()
	s.Equal(count.Int64(),int64(0))
	s.Nil(err)
}

package test

import (
	"bufio"
	"fmt"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/archaeologist"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"log"
	"math/big"
	"os/exec"
	"strings"
	"testing"
)

const RESURRECTION_TIME = int64(20)

type ArchTestSuite struct {
	suite.Suite
	config         *models.Config
	embalmerConfig *embalmer.EmbalmerConfig
	arch           *models.Archaeologist
	embalmer       *embalmer.Embalmer
	contractPort   string
	contractDir    string
}

func TestRunArchTestSuite(t *testing.T) {
	suite.Run(t, new(ArchTestSuite))
}

func (s *ArchTestSuite) exitContract() {
	args := []string{"./exit_contract.sh", s.contractPort}
	cmd := exec.Command("/bin/sh", args...)
	cmd.Start()
	cmd.Wait()
}

func (s *ArchTestSuite) deployContract() {
	args := []string{"./deploy_contract.sh", s.contractDir}
	cmd := exec.Command("/bin/sh", args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	readerLog := make(chan string)
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			readerLog <- scanner.Text()
		}
	}()
	cmd.Start()
	s.T().Log("Deploying Sarcophagus Contract...")
	/* Wait for contract to be deployed */
	for {
		select {
		case logLine := <-readerLog:
			if logLine == "name:    Sarcophagus" {
				return
			}
		}
	}
}

func (s *ArchTestSuite) deployArweave() {
	args := []string{"./deploy_arweave.sh", "kSyg2ajZbqiAE25AnkQcxHoGOnM5jUgT4t9TyZZQI3I"}
	cmd := exec.Command("/bin/sh", args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	readerLog := make(chan string)
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			readerLog <- scanner.Text()
		}
	}()
	cmd.Start()
	s.T().Log("Deploying Arweave...")
	/* Wait for contract to be deployed */
	for {
		select {
		case logLine := <-readerLog:
			if strings.Contains(logLine, "bridge") {
				return
			}
		}
	}
}

func (s *ArchTestSuite) initEnv() {
	viper.SetConfigName("test_setup")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find env file. It should be setup under test/.env")
		} else {
			log.Fatalf("Could not read env file. Please check it is configured correctly. Error: %v \n", err)
		}
	}
	s.contractPort = viper.GetString("CONTRACT_PORT")
	s.contractDir = viper.GetString("CONTRACT_DIR")
}

func (s *ArchTestSuite) SetupSuite() {
	config := new(models.Config)
	config.LoadConfig("test_config", "./", false)
	s.config = config
}

func (s *ArchTestSuite) SetupTest() {
	/* Start each test with a blank archaeologist and re-deployed contract */
	s.arch = new(models.Archaeologist)
	s.initEnv()
	s.exitContract()
	s.deployContract()
	s.deployArweave()
}

func (s *ArchTestSuite) AfterTest() {
	s.exitContract()
}

func (s *ArchTestSuite) InitEmbalmer() {
	config := new(embalmer.EmbalmerConfig)
	config.LoadEmbalmerConfig("embalmer_config", "../embalmer")
	s.embalmerConfig = config
	emb := new(embalmer.Embalmer)
	embalmer.InitEmbalmer(emb, config, RESURRECTION_TIME)
	s.embalmer = emb
	s.TransferSarcoToEmbalmer(big.NewInt(2000))
}

func (s *ArchTestSuite) TransferSarcoToEmbalmer(amount *big.Int) {
	_, err := s.arch.TokenSession.Transfer(s.embalmer.EmbalmerAddress, amount)
	if err != nil {
		s.T().Logf("ERROR TRANSFERRING SARCO TO EMBALMER: %v", err)
	}
}

func (s *ArchTestSuite) TestArchaeologistHappyPathWorkflow() {
	/* Archaeologist Initialize Without Errors */
	errStrings := archaeologist.InitializeArchaeologist(s.arch, s.config, ".")
	if len(errStrings) > 0 {
		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
	}
	s.Equal(0, len(errStrings))

	/* Archaeologist Registers with correct free bond amount*/
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	count, err := s.arch.SarcoSession.ArchaeologistCount()
	s.Equal(int64(1), count.Int64())
	s.Nil(err)

	contractArch, err := s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	s.Equal(int64(1000), contractArch.FreeBond.Int64())
	s.Equal(int64(0), s.arch.FreeBond.Int64())

	/* Archaeologist Updates with added free bond amount */
	s.arch.FreeBond = big.NewInt(500)
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	contractArch, err = s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	s.Equal(int64(1500), contractArch.FreeBond.Int64())

	/* Archaeologist Withdraws Free Bond Amount */
	s.arch.FreeBond = big.NewInt(-200)
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	contractArch, err = s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	s.Equal(int64(1300), contractArch.FreeBond.Int64())

	go archaeologist.EventsSubscribe(s.arch)

	/* Embalmer Creates First Sarco */
	s.InitEmbalmer()
	_, assetDoubleHashBytes := embalmer.DoubleHashBytesFromSeed(200, 20)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes, "Test Sarco")
	sarco, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
	s.Nil(err)
	s.Equal("Test Sarco", sarco.Name)
	s.Equal(1, len(s.arch.FileHandlers))
	s.Equal(1, len(s.arch.Sarcophaguses))
	s.Equal(sarco.ResurrectionTime, s.arch.Sarcophaguses[assetDoubleHashBytes])

	/* Embalmer Creates Second Sarco */
	_, assetDoubleHashBytesTwo := embalmer.DoubleHashBytesFromSeed(201, 20)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesTwo, "Test Sarco Two")
	sarcoTwo, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytesTwo)
	s.Nil(err)
	s.Equal("Test Sarco Two", sarcoTwo.Name)
	s.Equal(2, len(s.arch.FileHandlers))
	s.Equal(2, len(s.arch.Sarcophaguses))
	s.Equal(sarcoTwo.ResurrectionTime, s.arch.Sarcophaguses[assetDoubleHashBytesTwo])
}

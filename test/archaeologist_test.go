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

func (s *ArchTestSuite) SetupTest() {
	config := new(models.Config)
	config.LoadConfig("test_config", "./", false)
	s.config = config
	s.arch = new(models.Archaeologist)
	s.initEnv()
	s.exitContract()
	s.deployContract()
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
}

func (s *ArchTestSuite) TestRegisterArchaeologist() {
	errStrings := archaeologist.InitializeArchaeologist(s.arch, s.config, ".")
	if len(errStrings) > 0 {
		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
	}
	s.Equal(0, len(errStrings))

	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	count, err := s.arch.SarcoSession.ArchaeologistCount()
	s.Equal(int64(1), count.Int64())
	s.Nil(err)

	contractArch, err := s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	s.Equal(int64(1000), contractArch.FreeBond.Int64())
	s.Equal(int64(0), s.arch.FreeBond.Int64())

	s.arch.FreeBond = big.NewInt(500)

	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	contractArch, err = s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	s.Equal(int64(1500), contractArch.FreeBond.Int64())
	go archaeologist.EventsSubscribe(s.arch)

	s.InitEmbalmer()
	_, assetDoubleHashBytes := embalmer.DoubleHashBytesFromSeed(200, 20)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes, "Test Sarco")
	sarco, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
	s.Nil(err)
	s.Equal("Test Sarco", sarco.Name)
}

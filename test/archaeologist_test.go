package test

import (
	"bufio"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/archaeologist"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"log"
	"math/big"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	RESURRECTION_TIME = int64(15) // expressed in seconds from now
	EMBALMER_BALANCE = "4000000000000000000000" // 4000 sarco tokens
	FILE_BYTE_COUNT = 20 // size of single-encrypted file in bytes
)

type ArchTestSuite struct {
	suite.Suite
	config         *models.Config
	embalmerConfig *embalmer.EmbalmerConfig
	arch           *models.Archaeologist
	embalmer       *embalmer.Embalmer
	contractPort   string
	arweavePort    string
	contractDir    string
}

func TestRunArchTestSuite(t *testing.T) {
	suite.Run(t, new(ArchTestSuite))
}

func (s *ArchTestSuite) exitContract() {
	s.T().Logf("Exiting Sarco contract on port: %v", s.contractPort)
	args := []string{"./exit_contract.sh", s.contractPort}
	cmd := exec.Command("/bin/sh", args...)
	cmd.Start()
	cmd.Wait()
	time.Sleep(1 * time.Second)
}

func (s *ArchTestSuite) exitArweave() {
	args := []string{"./exit_arweave.sh", s.arweavePort}
	cmd := exec.Command("/bin/sh", args...)
	cmd.Start()
	cmd.Wait()
	time.Sleep(1 * time.Second)
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
			log.Print(logLine)
			if logLine == "name:    Sarcophagus" {
				return
			}
		}
	}
}

func (s *ArchTestSuite) deployArweave() {
	arweaveWallet, err := ar.InitArweaveWallet(s.config.ARWEAVE_KEY_FILE)
	if err != nil {
		s.T().Fatalf("Arweave Wallet could not be initialized: %v", err)
	}

	arweaveAddress := arweaveWallet.Address()

	s.T().Logf("Arweave Address: %v", arweaveAddress)

	args := []string{"./deploy_arweave.sh", s.arweavePort}
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
				args := []string{"./transfer_arweave.sh", arweaveAddress, s.arweavePort}
				cmd = exec.Command("/bin/sh", args...)
				cmd.Start()
				cmd.Wait()
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
	s.arweavePort = viper.GetString("ARWEAVE_PORT")
	s.contractDir = viper.GetString("CONTRACT_DIR")
}

func (s *ArchTestSuite) SetupSuite() {
	ecies.AddParamsForCurve(btcec.S256(), ecies.ECIES_AES128_SHA256)
	config := new(models.Config)
	config.LoadConfig("test_config", "./", false)
	s.config = config
}

func (s *ArchTestSuite) SetupTest() {
	/* Start each test with a blank archaeologist and re-deployed contract */
	s.T().Log("*** Setting up Test ***")

	s.arch = new(models.Archaeologist)
	s.initEnv()
	// s.exitContract()
	s.exitArweave()
	// s.deployContract()
	s.deployArweave()
}

func (s *ArchTestSuite) TeardownSuite() {
	s.T().Log("*** Stopping blockchains ***")

	// s.exitContract()
	s.exitArweave()
}

func (s *ArchTestSuite) InitEmbalmer() {
	config := new(embalmer.EmbalmerConfig)
	config.LoadEmbalmerConfig("embalmer_config", "../embalmer")
	s.embalmerConfig = config
	emb := new(embalmer.Embalmer)
	embalmer.InitEmbalmer(emb, config, RESURRECTION_TIME)
	s.embalmer = emb
	transferAmount, _ := new(big.Int).SetString(EMBALMER_BALANCE, 10)
	s.TransferSarcoToEmbalmer(transferAmount)
}

func (s *ArchTestSuite) simulateServiceRestart() {
	s.T().Log("Simulating Service Restart...")
	s.arch.FileHandlers = map[[32]byte]*big.Int{}
	s.arch.Sarcophaguses = map[[32]byte]*big.Int{}
	_ = archaeologist.InitializeArchaeologist(s.arch, s.config)
}

func (s *ArchTestSuite) TransferSarcoToEmbalmer(amount *big.Int) {
	_, err := s.arch.TokenSession.Transfer(s.embalmer.EmbalmerAddress, amount)
	if err != nil {
		s.T().Logf("ERROR TRANSFERRING SARCO TO EMBALMER: %v", err)
	}
}

//func (s *ArchTestSuite) TestTwoSarcosOneUnwrapTime() {
//	errStrings := archaeologist.InitializeArchaeologist(s.arch, s.config, ".")
//	if len(errStrings) > 0 {
//		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
//	}
//	s.Equal(0, len(errStrings))
//
//	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
//
//	go archaeologist.EventsSubscribe(s.arch)
//
//	s.InitEmbalmer()
//	timeUntilUnwrap := time.Until(time.Unix(s.embalmer.ResurrectionTime.Int64(), 0))
//
//	/* Sarco One Scheduled */
//	fileSeed := 200
//	fileBytes, assetDoubleHashBytes := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
//	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes, "Test Sarco")
//	s.embalmer.UpdateSarcophagus(assetDoubleHashBytes, fileBytes)
//	time.Sleep(3000 * time.Millisecond)
//
//	/* Sarco Two Scheduled */
//	fileSeed = 201
//	fileBytesTwo, assetDoubleHashBytesTwo := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
//	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesTwo, "Test Sarco Two")
//	s.embalmer.UpdateSarcophagus(assetDoubleHashBytesTwo, fileBytesTwo)
//
//	/* Wait for both sarcs to be unwrapped */
//	time.Sleep(timeUntilUnwrap)
//	time.Sleep(3000 * time.Millisecond)
//
//	/* Both sarcs should have a state of done */
//	sarcoOne, _ := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
//	s.Equal(uint8(2), sarcoOne.State)
//
//	sarcoTwo, _ := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
//	s.Equal(uint8(2), sarcoTwo.State)
//}

func (s *ArchTestSuite) TestArchaeologistHappyPathWorkflow() {
	/* Archaeologist Initializes Without Errors */
	errStrings := archaeologist.InitializeArchaeologist(s.arch, s.config)
	if len(errStrings) > 0 {
		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
	}
	s.Equal(0, len(errStrings))

	/* Archaeologist Registers with correct free bond amount */
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	count, err := s.arch.SarcoSession.ArchaeologistCount()
	s.Equal(int64(1), count.Int64())
	s.Nil(err)

	contractArch, err := s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	expectedVal, _ := new(big.Int).SetString("1000000000000000000000", 10)
	if expectedVal.Cmp(contractArch.FreeBond) != 0 {
		s.T().Log("Free bond should equal 1000000000000000000000")
		s.T().Fail()
	}
	s.Equal(int64(0), s.arch.FreeBond.Int64())

	/* Archaeologist Updates with added free bond amount */
	s.arch.FreeBond, _ = new(big.Int).SetString("500000000000000000000", 10)
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	contractArch, err = s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	expectedVal, _ = new(big.Int).SetString("1500000000000000000000", 10)
	if expectedVal.Cmp(contractArch.FreeBond) != 0 {
		s.T().Logf("Free bond should equal 1500000000000000000000 equals: %v", contractArch.FreeBond)
		s.T().Fail()
	}

	/* Archaeologist Withdraws Free Bond Amount */
	s.arch.FreeBond, _ = new(big.Int).SetString("-200000000000000000000", 10)
	archaeologist.RegisterOrUpdateArchaeologist(s.arch)
	contractArch, err = s.arch.SarcoSession.Archaeologists(s.arch.ArchAddress)
	expectedVal, _ = new(big.Int).SetString("1300000000000000000000", 10)
	if expectedVal.Cmp(contractArch.FreeBond) != 0 {
		s.T().Log("Free bond should equal 1300000000000000000000")
		s.T().Fail()
	}

	go archaeologist.EventsSubscribe(s.arch)

	/* Embalmer Creates First Sarco */
	/* Embalmer inits with balance of 4000 Sarco Tokens */
	s.InitEmbalmer()
	fileSeed := 200
	fileBytes, assetDoubleHashBytes := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes, "Test Sarco")
	sarco, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
	s.Nil(err)
	s.Equal("Test Sarco", sarco.Name)
	s.Equal(1, len(s.arch.FileHandlers))
	s.Equal(1, len(s.arch.Sarcophaguses))
	s.Equal(sarco.ResurrectionTime, s.arch.Sarcophaguses[assetDoubleHashBytes])

	/* Embalmer Creates Second Sarco */
	fileSeed += 1
	_, assetDoubleHashBytesTwo := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesTwo, "Test Sarco Two")
	sarcoTwo, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytesTwo)
	s.Nil(err)
	s.Equal("Test Sarco Two", sarcoTwo.Name)
	s.Equal(2, len(s.arch.FileHandlers))
	s.Equal(2, len(s.arch.Sarcophaguses))
	s.Equal(sarcoTwo.ResurrectionTime, s.arch.Sarcophaguses[assetDoubleHashBytesTwo])

	/* Embalmer Updates First Sarco */
	s.embalmer.UpdateSarcophagus(assetDoubleHashBytes, fileBytes)
	time.Sleep(4000 * time.Millisecond)
	s.Equal(2, len(s.arch.Sarcophaguses))
	s.Equal(1, len(s.arch.FileHandlers))
	s.Equal(1, s.arch.AccountIndex)

	/* Wait for unwrap and test unwrap result */
	/* State = 2 means sarco is 'done' */
	timeUntilUnwrap := time.Until(time.Unix(sarco.ResurrectionTime.Int64(), 0))
	time.Sleep(timeUntilUnwrap)
	time.Sleep(2000 * time.Millisecond)
	sarcoUnwrapped, err := s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytes)
	s.Nil(err)
	s.Equal(uint8(2), sarcoUnwrapped.State)
	s.Equal(1, len(s.arch.Sarcophaguses))

	/* Check state is correct on service restart */
	s.simulateServiceRestart()
	s.Equal(0, len(s.arch.Sarcophaguses))
	s.Equal(0, len(s.arch.FileHandlers))

	/* Embalmer Creates Third Sarco */
	s.embalmer.ResurrectionTime = big.NewInt(time.Now().Unix() + RESURRECTION_TIME)
	fileSeed += 1
	_, assetDoubleHashBytesThree := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesThree, "Test Sarco Three")

	/* Embalmer Creates Fourth Sarco */
	fileSeed += 1
	fileBytesFour, assetDoubleHashBytesFour := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesFour, "Test Sarco Four")

	/* Check state is correct on service restart */
	s.simulateServiceRestart()
	s.Equal(2, len(s.arch.Sarcophaguses))
	s.Equal(2, len(s.arch.FileHandlers))

	/* Embalmer Updates Fourth Sarco */
	s.embalmer.UpdateSarcophagus(assetDoubleHashBytesFour, fileBytesFour)
	time.Sleep(5000 * time.Millisecond)

	/* Check state is correct on service restart */
	s.simulateServiceRestart()
	s.Equal(1, len(s.arch.Sarcophaguses))
	s.Equal(s.embalmer.ResurrectionTime, s.arch.Sarcophaguses[assetDoubleHashBytesFour])
	s.Equal(0, len(s.arch.FileHandlers))

	/*
		Wait for unwrap and test unwrap result
		Because we just simulated a service restart, this will result in a duplicate unwrap event, as the previous unwrapping event for this sarc is not cancelled
		This would not happen if the service *actually* restarted b/c the previous timers would all be cancelled
		However, this is a good case to test.

		The nonce will be incorrect b/c unwrap event being called at the same time.
		The first unwrap should succeed, and the second should get rescheduled due to failure.
		The rescehduled unwrap should then detect the sarco state is done and just output a log and not be rescheduled again.
	*/

	timeUntilUnwrap = time.Until(time.Unix(s.embalmer.ResurrectionTime.Int64(), 0))
	time.Sleep(timeUntilUnwrap)
	time.Sleep(2000 * time.Millisecond)
	sarcoUnwrapped, err = s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytesFour)
	s.Nil(err)
	s.Equal(uint8(2), sarcoUnwrapped.State)
	s.Equal(0, len(s.arch.Sarcophaguses))
	time.Sleep(5000 * time.Millisecond)

	/* Embalmer Creates Fifth Sarco */
	s.embalmer.ResurrectionTime = big.NewInt(time.Now().Unix() + RESURRECTION_TIME)
	fileSeed += 1
	fileBytesFive, assetDoubleHashBytesFive := embalmer.DoubleHashBytesFromSeed(int64(fileSeed), FILE_BYTE_COUNT)
	s.embalmer.CreateSarcophagus(s.embalmerConfig.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytesFive, "Test Sarco Five")

	/* Embalmer Updates Fifth Sarco */
	s.embalmer.UpdateSarcophagus(assetDoubleHashBytesFive, fileBytesFive)

	/* Embalmer Rewraps Sarco for time < current resurrection time */
	s.embalmer.ResurrectionTime = big.NewInt(time.Now().Unix() + RESURRECTION_TIME - 10)
	s.embalmer.RewrapSarcophagus(assetDoubleHashBytesFive, s.embalmer.ResurrectionTime)

	/* Embalmer Rewraps Sarco for time > current resurrection time */
	s.embalmer.ResurrectionTime = big.NewInt(time.Now().Unix() + RESURRECTION_TIME)
	s.embalmer.RewrapSarcophagus(assetDoubleHashBytesFive, s.embalmer.ResurrectionTime)
	time.Sleep(10000 * time.Millisecond)

	/* State is still 1, original resurrection time did not cause an unwrap */
	sarcoUnwrapped, err = s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytesFive)
	s.Nil(err)
	s.Equal(uint8(1), sarcoUnwrapped.State)

	time.Sleep(5000 * time.Millisecond)

	/* Sarco is unwrapped after the final rewrap time */
	sarcoUnwrapped, err = s.arch.SarcoSession.Sarcophagus(assetDoubleHashBytesFive)
	s.Nil(err)
	s.Equal(uint8(2), sarcoUnwrapped.State)
	s.Equal(0, len(s.arch.Sarcophaguses))
	time.Sleep(8000 * time.Millisecond)
}

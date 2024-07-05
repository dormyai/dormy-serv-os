package listening

import (
	"context"
	"dormy/app/model"
	"dormy/config"
	"dormy/database"
	ml "dormy/middleware"
	"dormy/util"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// 监听dormy
func ListenToDormy() {
	client, err := ethclient.Dial(config.Get().Chains.MainWss)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// db := database.DB

	// db.Transaction(func(tx *gorm.DB) error {
	// 	processBlock(client, big.NewInt(44868411), tx)
	// 	return nil
	// })

	// 尝试读取最后处理的区块号
	lastProcessedBlock, err := getLastProcessedBlock()
	if err != nil {
		log.Fatalf("Failed to retrieve the last processed block: %v", err)
	}

	if lastProcessedBlock.Int64() > 0 { //大于0才处理
		go dealOldBlock(client, lastProcessedBlock)
	}

	// 从当前区块开始实时监听新区块
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		ml.Log.Infof("Failed to subscribe to new headers: %v", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			ml.Log.Info(err)
		case header := <-headers:

			db := database.DB
			db.Transaction(func(tx *gorm.DB) error {

				err := processBlock(client, header.Number, tx)

				if err != nil {
					return err
				}

				err = saveLastProcessedBlock(header.Number, tx)

				return err
			})

		}
	}
}

func dealOldBlock(client *ethclient.Client, lastProcessedBlock *big.Int) {
	// 处理从最后一个已处理区块到当前区块的所有区块
	currentBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve the current block: %v", err)
	}

	// 补区块
	for i := new(big.Int).Set(lastProcessedBlock); i.Cmp(currentBlock.Number()) <= 0; i.Add(i, big.NewInt(1)) {

		db := database.DB

		db.Transaction(func(tx *gorm.DB) error {

			err := processBlock(client, i, tx)

			if err != nil {
				return err
			}

			err = saveLastProcessedBlock(i, tx)

			return err

		})

	}
}

// 获取到最后一个处理的区块高度
func getLastProcessedBlock() (*big.Int, error) {
	db := database.DB
	var lastBlock model.ProcessedBlock
	db.Where(" status = -1 ").Order("block_number desc").First(&lastBlock)

	if lastBlock != (model.ProcessedBlock{}) {

		return big.NewInt(lastBlock.BlockNumber), nil
	}

	return big.NewInt(0), nil
}

func saveLastProcessedBlock(blockNumber *big.Int, tx *gorm.DB) error {
	db := database.DB

	var lastBlock model.ProcessedBlock
	lastBlock.ID, _ = util.Sf.GenerateID()
	lastBlock.BlockNumber = blockNumber.Int64()
	lastBlock.CreateAt = time.Now()
	lastBlock.Status = 0
	r := db.Create(lastBlock)

	return r.Error
}

func processBlock(client *ethclient.Client, blockNumber *big.Int, tx *gorm.DB) error {
	// 这里可以添加处理区块的逻辑，例如检索区块中的日志
	// fmt.Println("Processing block:", blockNumber)

	// client, err := ethclient.Dial(config.Get().Chains.Main)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// 	return err
	// }

	contractAddress := common.HexToAddress(config.Get().Dormy.Dormy)

	eventSignMint := []byte("MintDormy(address,uint256,uint256,uint256)") //mint 合约
	eventSignMintHash := crypto.Keccak256Hash(eventSignMint)

	// fmt.Println("eventSignMintHash:", eventSignMintHash)

	// 指定区块的范围（这里以单个区块为例）
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventSignMintHash}},
	}

	// 检索日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to retrieve logs: %v", err)
	}

	// 分析日志
	for _, vLog := range logs {
		fmt.Printf("Log found in block number %d\n", vLog.BlockNumber)

		jsonData, err := json.Marshal(vLog)
		if err != nil {
			log.Fatalf("Error marshalling vLog to JSON: %v", err)
		}
		ml.Log.Infof("Data: %s\n", jsonData) // 事件的数据部分

		//mint 事件
		if vLog.Topics[0].Hex() == "0xabd23447544f9a441454c0ae1081cc91194dec5c91e798456bad2819756ce4c7" {

			encodedAddress := vLog.Topics[1]
			toAddress := "0x" + string(encodedAddress.String())[26:]

			ml.Log.Info("toAddress:", toAddress)

			// ml.Log.Info("solt:", string(vLog.Topics[2].String()[2:]))
			// solt, err := strconv.ParseInt(string(vLog.Topics[2].String()[2:]), 10, 64)
			// if err != nil {
			// 	fmt.Println("转换出错1:", err)
			// 	continue
			// }

			tokenId := new(big.Int).SetBytes(vLog.Topics[3].Bytes())

			solt := new(big.Int).SetBytes(vLog.Topics[2].Bytes())

			amount := new(big.Int).SetBytes(vLog.Data)

			processingMint(eventSignMintHash.String(), vLog.TxHash.Hex(), toAddress, solt.Int64(), amount.Int64(), tokenId.Int64())
		}

	}

	return nil
}

func processingMint(topicHash string, transactionHash string, to string, solt int64, amount int64, tokenId int64) error {
	ml.Log.Infof("processingMint: %s %s %s %d %d ", topicHash, transactionHash, to, solt, amount)

	var propertyChainInfo model.PropertyChainInfo
	db := database.DB
	db.Where(" solt_id = ? ", solt).First(&propertyChainInfo)

	var user model.AccountUserInfo
	db.Where(" upper(address) = upper(?) ", to).First(&user)

	var propertyUserSummary model.PropertyUserSummary

	propertyUserSummary.PropertyInfoID = propertyChainInfo.PropertyInfoID
	propertyUserSummary.Address = to
	propertyUserSummary.UserID = user.ID
	propertyUserSummary.Amount = amount
	propertyUserSummary.StartTime = time.Now()
	propertyUserSummary.Status = 1 //有效 持有中
	propertyUserSummary.TxHash = transactionHash
	propertyUserSummary.EventSignHash = topicHash
	propertyUserSummary.TokenID = tokenId

	r := db.Create(&propertyUserSummary)

	return r.Error
}

package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		MongoDB
		Chain
		Contract
		Secret
	}
	MongoDB struct {
		DatabaseName string `env:"MONGO_DATABASE_NAME"`
		URLString    string `env:"MONGO_URL_STRING"`
	}
	Chain struct {
		ID   int64  `env:"CHAIN_ID"`
		Name string `env:"CHAIN_NAME"`
		RPC  string `env:"CHAIN_RPC"`
	}
	Contract struct {
		Address string `env:"CONTRACT_ADDRESS"`
	}
	Secret struct {
		MasterWalletPrivateKey string `env:"MASTER_WALLET_PRIVATE_KEY"`
	}
)

var C Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Application config parsing failed: " + err.Error() + " => Exit!")
		return
	}

	cfg := &C
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalln("Application config parsing failed: " + err.Error() + " => Exit!")
		return
	}

	log.Println("Load Config Successfully!")
}

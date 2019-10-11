package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MerchantConfig struct {
	Servers          servers
	Block_nodes      blockNodes
	Wallet_key_rsa   walletKeyRsa
	Import_src_addrs importCoins
	Servers_key      serversKey
	Sign             sign
}

type servers struct {
	Account_url   string
	User_url      string
	Yyt_agent_url string
	Notice_url    string
}

type serversKey struct {
	Yyt_agent_key string
}

type blockNodes struct {
	Btc_nodes []btcBlockNode
}

type btcBlockNode struct {
	Coin string
	Host string
	Port int
	User string
	Pwd  string
}

type walletKeyRsa struct {
	Pri_key string
	Pub_key string
}

type importCoins struct {
	Coins []importCoin
}

type importCoin struct {
	Coin string
	Addr string
}

type sign struct {
	PriSign string
	PubSign string
}

var Config MerchantConfig

func init() {
	New("./config/config.yml")
	log.Printf("server配置:%+v\n", Config.Block_nodes)
}

func New(path string) MerchantConfig {
	data, _ := ioutil.ReadFile(path)
	Config = MerchantConfig{}
	yaml.Unmarshal(data, &Config)
	log.Printf("server配置:%+v\n", Config)
	return Config
}

func (c MerchantConfig) BtcBlockNodes() []btcBlockNode {
	return c.Block_nodes.Btc_nodes
}

func (c MerchantConfig) ExistsBlockNode(coin string) bool {
	nodes := Config.BtcBlockNodes()
	for _, node := range nodes {
		if node.Coin == coin {
			return true
		}
	}
	return false
}

func (c MerchantConfig) ImportAddr(coin string) string {
	for _, importCoin := range c.Import_src_addrs.Coins {
		if importCoin.Coin == coin {
			return importCoin.Addr
		}
	}
	return ""
}

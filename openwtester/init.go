package openwtester

import (
	"github.com/blocktree/digibyte-adapter/digibyte"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(digibyte.Symbol, digibyte.NewWalletManager())
}

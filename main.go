package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0fatih/supher/disasm"
	"github.com/0fatih/supher/utils"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	file := flag.String("file", "", "Disassemble from a bin file")
	tx := flag.String("tx", "", "DisAssemble from a transaction")
	network := flag.String("network", "ethereum", "Network that you want to process")
	flag.Parse()

	if len(*file) > 0 {
		code, err := utils.ReadFromFile(*file)
		if err != nil {
			utils.ColoredPrint("[!] failed to read from file!", utils.PrintColors.RED)
			os.Exit(1)
		}

		disasm.ListenForDisasm(code)
	}

	if len(*tx) > 0 {
		isConfigFileExists := utils.CheckConfigFile()
		if isConfigFileExists != true {
			err := utils.CreateConfigFile()
			if err != nil {
				utils.ColoredPrint("[!] error occured while creating config file:"+err.Error(), utils.PrintColors.RED)
				os.Exit(1)
			}

			utils.ColoredPrint("[~] please edit your config file in"+utils.CONFIG_FILE, utils.PrintColors.YELLOW)
			os.Exit(1)
		}

		isNetworkSettingsCorrect := utils.CheckNetworkSettings(*network)
		if isNetworkSettingsCorrect != true {
			utils.ColoredPrint("[~] please edit your config file in"+utils.CONFIG_FILE, utils.PrintColors.YELLOW)
			os.Exit(1)
		}

		txData, err := utils.GetTransactionData(common.HexToHash(*tx), *network)
		if err != nil {
			utils.ColoredPrint("[!] error eccoured while getting transaction data", utils.PrintColors.RED)
			os.Exit(1)
		}

		fmt.Println(common.Bytes2Hex(txData))
		disasm.ListenForDisasm(txData)
	}
}

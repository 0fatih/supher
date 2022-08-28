package disasm

import (
	"fmt"

	"github.com/0fatih/supher/utils"
	"github.com/eiannone/keyboard"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func ConvertAndDisAssemble(codeStr string) {
	code := []uint8(codeStr)
	ListenForDisasm(code)
}

func ListenForDisasm(code []uint8) {
	utils.ColoredPrint("[*] Press N for next opcode, ESC for close the program...", "green")

	// Listen keyboard
	keysEvents, err := keyboard.GetKeys(1)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	DisAssemble(code, keysEvents)
}

func DisAssemble(code []uint8, keysEvents <-chan keyboard.KeyEvent) {
	code = common.Hex2Bytes(string(code[:len(code)-1]))

	for pc := uint64(0); pc < uint64(len(code)); pc++ {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeyEsc {
			break
		}
		if event.Rune != 'n' {
			continue
		}

		op := vm.OpCode(code[pc])

		toPrint := fmt.Sprintf("%-5d %v", pc, op)
		utils.ColoredPrintf(toPrint, utils.PrintColors.CYAN)

		switch op {
		case vm.PUSH1, vm.PUSH2, vm.PUSH3, vm.PUSH4, vm.PUSH5, vm.PUSH6, vm.PUSH7, vm.PUSH8, vm.PUSH9, vm.PUSH10, vm.PUSH11, vm.PUSH12, vm.PUSH13, vm.PUSH14, vm.PUSH15, vm.PUSH16, vm.PUSH17, vm.PUSH18, vm.PUSH19, vm.PUSH20, vm.PUSH21, vm.PUSH22, vm.PUSH23, vm.PUSH24, vm.PUSH25, vm.PUSH26, vm.PUSH27, vm.PUSH28, vm.PUSH29, vm.PUSH30, vm.PUSH31, vm.PUSH32:
			a := uint64(op) - uint64(vm.PUSH1) + 1

			toPrint = fmt.Sprintf("	=> %x", code[pc+1:pc+1+a])
			utils.ColoredPrintf(toPrint, utils.PrintColors.CYAN)

			pc += a
		}
		fmt.Println()
	}

}

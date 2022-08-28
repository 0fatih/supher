# Supher

Currently this program only disassembly bytecode to assembly code for Solidity. 
But in the future I want to add some features like:

- Disassembly directly `.sol` files.
- Debug transactions with only transaction hash.
- Add visual supports (stack, memory etc.).
- Deploy contracts.
- Get codes of the verified smart contracts (using etherscan-like apis).
- Send transactions using flags (i.e, `--contract 0x0 --transfer --from 15 --to 20`)

# Usage

  ```
  supher --file erc20Token.bin
  ```

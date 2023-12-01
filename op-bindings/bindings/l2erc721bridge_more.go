// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_uint256)49_storage\"}],\"types\":{\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100885760003560e01c80637f46ddb21161005b5780637f46ddb21461014a578063927ede2d14610171578063aa55745214610198578063c89701a2146101ab57600080fd5b80633687011a1461008d5780633cb747bf146100a257806354fd4d50146100ee578063761f449314610137575b600080fd5b6100a061009b36600461101c565b6101d1565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012a6040518060400160405280600581526020017f312e352e3000000000000000000000000000000000000000000000000000000081525081565b6040516100e5919061110a565b6100a061014536600461111d565b61027d565b6100c47f000000000000000000000000000000000000000000000000000000000000000081565b6100c47f000000000000000000000000000000000000000000000000000000000000000081565b6100a06101a63660046111b5565b6107e4565b7f00000000000000000000000000000000000000000000000000000000000000006100c4565b333b15610265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b61027586863333888888886108a0565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561039b57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610383919061122c565b73ffffffffffffffffffffffffffffffffffffffff16145b610427576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f746865722062726964676500606482015260840161025c565b3073ffffffffffffffffffffffffffffffffffffffff8816036104cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c6600000000000000000000000000000000000000000000606482015260840161025c565b6104f6877f74259ebf00000000000000000000000000000000000000000000000000000000610e3e565b610582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e7400000000000000000000606482015260840161025c565b8673ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f1919061122c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146106d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604b60248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433732312060648201527f6c6f63616c20746f6b656e000000000000000000000000000000000000000000608482015260a40161025c565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b15801561074157600080fd5b505af1158015610755573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107d39493929190611292565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516610887576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f74206265206164647265737328302900000000000000000000000000000000606482015260840161025c565b61089787873388888888886108a0565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff8716610943576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f742062652061646472657373283029000000000000000000000000000000606482015260840161025c565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa1580156109ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d2919061122c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610a8c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e65720000606482015260840161025c565b60008873ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ad9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afd919061122c565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610bba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c7565000000000000000000606482015260840161025c565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610c2a57600080fd5b505af1158015610c3e573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610c6997969594939291906112d2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610d7e907f00000000000000000000000000000000000000000000000000000000000000009085908a9060040161132f565b600060405180830381600087803b158015610d9857600080fd5b505af1158015610dac573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610e2a9493929190611292565b60405180910390a450505050505050505050565b6000610e4983610e61565b8015610e5a5750610e5a8383610ec6565b9392505050565b6000610e8d827f01ffc9a700000000000000000000000000000000000000000000000000000000610ec6565b8015610ec05750610ebe827fffffffff00000000000000000000000000000000000000000000000000000000610ec6565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015610f7e575060208210155b8015610f8a5750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610fb757600080fd5b50565b803563ffffffff81168114610fce57600080fd5b919050565b60008083601f840112610fe557600080fd5b50813567ffffffffffffffff811115610ffd57600080fd5b60208301915083602082850101111561101557600080fd5b9250929050565b60008060008060008060a0878903121561103557600080fd5b863561104081610f95565b9550602087013561105081610f95565b94506040870135935061106560608801610fba565b9250608087013567ffffffffffffffff81111561108157600080fd5b61108d89828a01610fd3565b979a9699509497509295939492505050565b6000815180845260005b818110156110c5576020818501810151868301820152016110a9565b818111156110d7576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e5a602083018461109f565b600080600080600080600060c0888a03121561113857600080fd5b873561114381610f95565b9650602088013561115381610f95565b9550604088013561116381610f95565b9450606088013561117381610f95565b93506080880135925060a088013567ffffffffffffffff81111561119657600080fd5b6111a28a828b01610fd3565b989b979a50959850939692959293505050565b600080600080600080600060c0888a0312156111d057600080fd5b87356111db81610f95565b965060208801356111eb81610f95565b955060408801356111fb81610f95565b94506060880135935061121060808901610fba565b925060a088013567ffffffffffffffff81111561119657600080fd5b60006020828403121561123e57600080fd5b8151610e5a81610f95565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006112c8606083018486611249565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261132260c083018486611249565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061135e606083018561109f565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
	immutableReferences["L2ERC721Bridge"] = true
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)47_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)47_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106101125760003560e01c80635c975abb116100a55780638f601f6611610074578063a3a7954811610059578063a3a795481461043e578063c89701a21461023f578063e11013dd1461045157600080fd5b80638f601f66146103c4578063927ede2d1461040a57600080fd5b80635c975abb14610341578063662a633a1461035d5780637f46ddb21461037057806387087623146103a457600080fd5b806336c717c1116100e157806336c717c11461023f5780633cb747bf14610298578063540abf73146102cb57806354fd4d50146102eb57600080fd5b80630166a07a146101e657806309fc8843146102065780631635f5fd1461021957806332b7006d1461022c57600080fd5b366101e157333b156101ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101df73deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d4060405180602001604052806000815250610464565b005b600080fd5b3480156101f257600080fd5b506101df6102013660046122a6565b61053f565b6101df610214366004612357565b61092c565b6101df6102273660046123aa565b610a03565b6101df61023a36600461241d565b610ed0565b34801561024b57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102a457600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061026e565b3480156102d757600080fd5b506101df6102e6366004612471565b610faa565b3480156102f757600080fd5b506103346040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161028f919061255e565b34801561034d57600080fd5b506040516000815260200161028f565b6101df61036b3660046122a6565b610fef565b34801561037c57600080fd5b5061026e7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103b057600080fd5b506101df6103bf366004612571565b611062565b3480156103d057600080fd5b506103fc6103df3660046125f4565b600260209081526000928352604080842090915290825290205481565b60405190815260200161028f565b34801561041657600080fd5b5061026e7f000000000000000000000000000000000000000000000000000000000000000081565b6101df61044c366004612571565b611136565b6101df61045f36600461262d565b61117a565b7fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff8716016104b3576104ae85858585856111c3565b610537565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610500573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105249190612690565b9050610535878288888888886113a7565b505b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561065d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610621573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106459190612690565b73ffffffffffffffffffffffffffffffffffffffff16145b61070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b610718876116ee565b15610866576107278787611750565b6107d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561084957600080fd5b505af115801561085d573d6000803e3d6000fd5b505050506108e8565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546108a49084906126dc565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c16835293905291909120919091556108e8908585611870565b610535878787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061194492505050565b333b156109bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b6109fe3333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506111c392505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610b2157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ae5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b099190612690565b73ffffffffffffffffffffffffffffffffffffffff16145b610bd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b823414610c62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101a2565b3073ffffffffffffffffffffffffffffffffffffffff851603610d07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101a2565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610de2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101a2565b610e2485858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506119d292505050565b6000610e41855a8660405180602001604052806000815250611a73565b905080610537576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101a2565b333b15610f5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b610fa3853333878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046492505050565b5050505050565b61053587873388888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a792505050565b73ffffffffffffffffffffffffffffffffffffffff871615801561103c575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b156110535761104e8585858585610a03565b610535565b6105358688878787878761053f565b333b156110f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b61053786863333888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a792505050565b610537863387878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046492505050565b6111bd3385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506111c392505050565b50505050565b823414611252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101a2565b61125e85858584611a8d565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016112db94939291906126f3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261136e9291889060040161273c565b6000604051808303818588803b15801561138757600080fd5b505af115801561139b573d6000803e3d6000fd5b50505050505050505050565b6113b0876116ee565b156114fe576113bf8787611750565b611471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156114e157600080fd5b505af11580156114f5573d6000803e3d6000fd5b50505050611592565b61152073ffffffffffffffffffffffffffffffffffffffff8816863086611b2e565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461155e908490612781565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b6115a0878787878786611b8c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b898b8a8a8a8960405160240161162096959493929190612799565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b90921682526116b39291879060040161273c565b600060405180830381600087803b1580156116cd57600080fd5b505af11580156116e1573d6000803e3d6000fd5b5050505050505050505050565b600061171a827f1d1d8b6300000000000000000000000000000000000000000000000000000000611c1a565b8061174a575061174a827fec4fc8e300000000000000000000000000000000000000000000000000000000611c1a565b92915050565b600061177c837f1d1d8b6300000000000000000000000000000000000000000000000000000000611c1a565b15611825578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f09190612690565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905061174a565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117cc573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109fe9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611c3d565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd898686866040516119bc939291906127f4565b60405180910390a4610537868686868686611d49565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611a5f939291906127f4565b60405180910390a46111bd84848484611dd1565b600080600080845160208601878a8af19695505050505050565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611b1a939291906127f4565b60405180910390a46111bd84848484611e3e565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526111bd9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016118c2565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611c04939291906127f4565b60405180910390a4610537868686868686611e9d565b6000611c2583611f15565b8015611c365750611c368383611f79565b9392505050565b6000611c9f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166120489092919063ffffffff16565b8051909150156109fe5780806020019051810190611cbd9190612832565b6109fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101a2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611dc1939291906127f4565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051611e30929190612854565b60405180910390a350505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051611e30929190612854565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611dc1939291906127f4565b6000611f41827f01ffc9a700000000000000000000000000000000000000000000000000000000611f79565b801561174a5750611f72827fffffffff00000000000000000000000000000000000000000000000000000000611f79565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612031575060208210155b801561203d5750600081115b979650505050505050565b6060612057848460008561205f565b949350505050565b6060824710156120f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101a2565b73ffffffffffffffffffffffffffffffffffffffff85163b61216f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101a2565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612198919061286d565b60006040518083038185875af1925050503d80600081146121d5576040519150601f19603f3d011682016040523d82523d6000602084013e6121da565b606091505b509150915061203d828286606083156121f4575081611c36565b8251156122045782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a2919061255e565b73ffffffffffffffffffffffffffffffffffffffff8116811461225a57600080fd5b50565b60008083601f84011261226f57600080fd5b50813567ffffffffffffffff81111561228757600080fd5b60208301915083602082850101111561229f57600080fd5b9250929050565b600080600080600080600060c0888a0312156122c157600080fd5b87356122cc81612238565b965060208801356122dc81612238565b955060408801356122ec81612238565b945060608801356122fc81612238565b93506080880135925060a088013567ffffffffffffffff81111561231f57600080fd5b61232b8a828b0161225d565b989b979a50959850939692959293505050565b803563ffffffff8116811461235257600080fd5b919050565b60008060006040848603121561236c57600080fd5b6123758461233e565b9250602084013567ffffffffffffffff81111561239157600080fd5b61239d8682870161225d565b9497909650939450505050565b6000806000806000608086880312156123c257600080fd5b85356123cd81612238565b945060208601356123dd81612238565b935060408601359250606086013567ffffffffffffffff81111561240057600080fd5b61240c8882890161225d565b969995985093965092949392505050565b60008060008060006080868803121561243557600080fd5b853561244081612238565b9450602086013593506124556040870161233e565b9250606086013567ffffffffffffffff81111561240057600080fd5b600080600080600080600060c0888a03121561248c57600080fd5b873561249781612238565b965060208801356124a781612238565b955060408801356124b781612238565b9450606088013593506124cc6080890161233e565b925060a088013567ffffffffffffffff81111561231f57600080fd5b60005b838110156125035781810151838201526020016124eb565b838111156111bd5750506000910152565b6000815180845261252c8160208601602086016124e8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611c366020830184612514565b60008060008060008060a0878903121561258a57600080fd5b863561259581612238565b955060208701356125a581612238565b9450604087013593506125ba6060880161233e565b9250608087013567ffffffffffffffff8111156125d657600080fd5b6125e289828a0161225d565b979a9699509497509295939492505050565b6000806040838503121561260757600080fd5b823561261281612238565b9150602083013561262281612238565b809150509250929050565b6000806000806060858703121561264357600080fd5b843561264e81612238565b935061265c6020860161233e565b9250604085013567ffffffffffffffff81111561267857600080fd5b6126848782880161225d565b95989497509550505050565b6000602082840312156126a257600080fd5b8151611c3681612238565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156126ee576126ee6126ad565b500390565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526127326080830184612514565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061276b6060830185612514565b905063ffffffff83166040830152949350505050565b60008219821115612794576127946126ad565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526127e860c0830184612514565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006128296060830184612514565b95945050505050565b60006020828403121561284457600080fd5b81518015158114611c3657600080fd5b8281526040602082015260006120576040830184612514565b6000825161287f8184602087016124e8565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
	immutableReferences["L2StandardBridge"] = true
}

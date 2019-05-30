package digibyte_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcdrivers/btcTransaction"
)

const (
	btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	MainNetAddressPrefix = btcTransaction.AddressPrefix{[]byte{0x1e}, []byte{0x3f}, "dgb"}
	TestNetAddressPrefix = btcTransaction.AddressPrefix{[]byte{0x7e}, []byte{0x8c}, "dgbt"}

	DGB_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x1e}, Suffix: nil}
	DGB_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x7e}, Suffix: nil}
	DGB_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x80}, Suffix: nil}
	DGB_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x80}, Suffix: nil}
	DGB_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x3f}, Suffix: nil}
	DGB_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x8c}, Suffix: nil}
	DGB_mainnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "dbg", "h160", 20, nil, nil}
	DGB_testnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "dbgt", "h160", 20, nil, nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := DGB_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = DGB_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := DGB_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = DGB_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}

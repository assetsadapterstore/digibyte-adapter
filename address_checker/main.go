package main

import (
	"C"
	"errors"

	"github.com/assetsadapterstore/digibyte-adapter/digibyte_addrdec"
	encoder "github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	ErrorSymbolType = errors.New("Invalid symbol type!!!")
)

// Check is valdating an address
func Check(address string) (bool, error) {

	decoded, err := encoder.Base58Decode(address, encoder.NewBase58Alphabet(digibyte_addrdec.BtcAlphabet))
	if err != nil {
		return false, err
	}

	if decoded[0] == digibyte_addrdec.DGB_mainnetAddressP2PKH.Prefix[0] {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_mainnetAddressP2PKH)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}
	if decoded[0] == digibyte_addrdec.DGB_mainnetAddressP2SH.Prefix[0] {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_mainnetAddressP2SH)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}
	if address[0] == 'd' && address[1] == 'g' && address[2] == 'b' {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_mainnetAddressBech32V0)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}
	if decoded[0] == digibyte_addrdec.DGB_testnetAddressP2PKH.Prefix[0] {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_testnetAddressP2PKH)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}
	if decoded[0] == digibyte_addrdec.DGB_testnetAddressP2SH.Prefix[0] {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_testnetAddressP2SH)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}
	if address[0] == 'd' && address[1] == 'g' && address[2] == 'b' && address[3] == 't' {
		_, err := encoder.AddressDecode(address, digibyte_addrdec.DGB_testnetAddressBech32V0)
		if err == nil {
			return true, err
		} else {
			return false, err
		}
	}

	return false, ErrorSymbolType

}

func main() {}

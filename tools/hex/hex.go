package hex

import (
	"fmt"
	"math/big"
)

func HexToBigInt(hexStr string) (*big.Int, error) {
	value := new(big.Int)

	_, success := value.SetString(hexStr, 0)
	if !success {
		return nil, fmt.Errorf("conversion not successfull")
	}

	return value, nil
}

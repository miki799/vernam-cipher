package utils

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func ReadTextFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}

func GetMessageBitsLength(message string) int {
	return len(message) * 8
}

func ConvertStringToBigInt(str string) []*big.Int {
	strBytes := []byte(str)

	bigInts := make([]*big.Int, 0, len(strBytes)*8)
	for _, b := range strBytes {
		for i := 0; i < 8; i++ {
			bit := int(b>>uint(7-i)) & 1
			bigInts = append(bigInts, big.NewInt(int64(bit)))
		}
	}

	return bigInts
}

func ConvertBigIntToString(bits []*big.Int) (string, error) {
	byteSlice := make([]byte, 0, (len(bits)+7)/8)

	var currentByte byte
	bitIndex := 7

	for _, bit := range bits {
		if bit.Sign() != 0 && bit.Sign() != 1 {
			return "", fmt.Errorf("invalid bit value: %s", bit.String())
		}

		bitValue := bit.Int64()

		currentByte |= byte(bitValue << uint(bitIndex))
		bitIndex--

		if bitIndex < 0 {
			byteSlice = append(byteSlice, currentByte)
			currentByte = 0
			bitIndex = 7
		}
	}

	return string(byteSlice), nil
}

package vernam

import (
	"fmt"
	"math/big"

	"github.com/miki799/vernam-cipher/bbs"
	"github.com/miki799/vernam-cipher/utils"
)

func CreateKey(bits int) ([]*big.Int, error) {

	bbs, err := bbs.BlumBlumShubGenerator(bits)
	if err != nil {
		return nil, err
	}

	key := make([]*big.Int, 0)
	for i := 0; i < bits; i++ {
		randomBit := bbs.GenerateNextBit()
		key = append(key, randomBit)
	}

	return key, nil
}

func Encrypt(message string, key []*big.Int) ([]*big.Int, error) {

	bits := utils.GetMessageBitsLength(message)
	messageBitsInt := utils.ConvertStringToBigInt(message)

	if len(messageBitsInt) != len(key) {
		return nil, fmt.Errorf("message and key are of different length")
	}

	cryptogram := make([]*big.Int, 0)

	for i := 0; i < bits; i++ {
		tmp := new(big.Int)
		tmp.Add(messageBitsInt[i], key[i])
		tmp.Mod(tmp, bbs.TWO)
		cryptogram = append(cryptogram, tmp)
	}

	return cryptogram, nil
}

func Decrypt(cryptogram, key []*big.Int) (string, error) {

	bits := len(cryptogram)

	if len(cryptogram) != len(key) {
		return "", fmt.Errorf("cryptogram and key are of different length")
	}

	plaintext := make([]*big.Int, 0)

	for i := 0; i < bits; i++ {
		tmp := new(big.Int)
		tmp.Add(cryptogram[i], key[i])
		tmp.Mod(tmp, bbs.TWO)
		plaintext = append(plaintext, tmp)
	}

	plaintextStr, err := utils.ConvertBigIntToString(plaintext)
	if err != nil {
		return "", err
	}

	return plaintextStr, nil
}

func Verify(originalMessage, encryptedMessage string) bool {
	return originalMessage == encryptedMessage
}

package bbs

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var ONE *big.Int = big.NewInt(1)
var TWO *big.Int = big.NewInt(2)

type BBS struct {
	p *big.Int
	q *big.Int
	n *big.Int
	x *big.Int
}

func (bbs *BBS) GenerateNextBit() *big.Int {
	bbs.x = new(big.Int).Exp(bbs.x, TWO, bbs.n)
	return new(big.Int).Mod(bbs.x, TWO)
}

func (bbs *BBS) String() string {
	return fmt.Sprintf("BBS {\n p: %v\n q: %v\n n: %v\n x: %v\n}", bbs.p, bbs.q, bbs.n, bbs.x)
}

func BlumBlumShubGenerator(bits int) (*BBS, error) {
	p, err1 := generatePrimeNumber(bits)
	if err1 != nil {
		return nil, err1
	}

	q, err2 := generatePrimeNumber(bits)
	if err2 != nil {
		return nil, err2
	}

	n := new(big.Int).Mul(p, q)

	seed, err3 := generateSeed(n)
	if err3 != nil {
		return nil, err3
	}

	return &BBS{p, q, n, seed}, nil
}

func generatePrimeNumber(bits int) (*big.Int, error) {
	for {
		num, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, fmt.Errorf("rand.Prime() returned error: %v", err)
		}

		if num.ProbablyPrime(bits) {
			return num, nil
		}

		fmt.Println("Generated number is not prime!", err)
	}
}

/*
Generates a random seed within the range [1, n-1]
*/
func generateSeed(n *big.Int) (*big.Int, error) {
	max := new(big.Int).Sub(n, ONE)

	for {
		seed, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, fmt.Errorf("rand.Prime() returned error: %v", err)
		}

		// Add 1 to ensure seed is in the range [1, n-1]
		seed.Add(seed, ONE)

		// Ensure seed is relatively prime to n
		if new(big.Int).GCD(nil, nil, seed, n).Cmp(ONE) == 0 {
			return seed, nil
		}
	}
}

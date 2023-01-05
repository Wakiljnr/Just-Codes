package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"time"
)

const maxNonce = math.MaxUint32 // 4 billion

func proofOfWork(header string, difficultyBits uint) (string, uint32) {
	// calculate the difficulty target
	target := math.Pow(2, float64(256-difficultyBits))

	for nonce := uint32(0); nonce < maxNonce; nonce++ {
		hashResult := sha256.Sum256([]byte(header + strconv.FormatUint(uint64(nonce), 10)))
		hashString := hex.EncodeToString(hashResult[:])

		// check if this is a valid result, equal to or below the target
		hashInt, _ := strconv.ParseUint(hashString, 16, 64)
		if float64(hashInt) <= target {
			fmt.Println(fmt.Sprintf("Success with nonce %d", nonce))
			fmt.Println(fmt.Sprintf("Hash is %s", hashString))
			return hashString, nonce
		}
	}

	fmt.Println(fmt.Sprintf("Failed after %d (maxNonce) tries", maxNonce))
	return "", 0
}

func main() {
	var nonce uint32
	var hashResult string

	// difficulty from 0 to 31 bits
	for difficultyBits := uint(0); difficultyBits < 32; difficultyBits++ {
		difficulty := uint64(math.Pow(2, float64(difficultyBits)))
		fmt.Println(fmt.Sprintf("Difficulty: %d (%d bits)", difficulty, difficultyBits))
		fmt.Println("Starting search...")

		// checkpoint the current time
		startTime := time.Now()

		// make a new block which includes the hash from the previous block
		// we fake a block of transactions - just a string
		newBlock := "test block with transactions" + hashResult

		// find a valid nonce for the new block
		hashResult, nonce = proofOfWork(newBlock, difficultyBits)

		// checkpoint how long it took to find a result
		elapsedTime := time.Since(startTime)
		fmt.Println(fmt.Sprintf("Elapsed Time: %s", elapsedTime))

		if elapsedTime > 0 {
			// estimate the hashes per second
			hashPower := uint64(nonce) / uint64(elapsedTime.Seconds())
			fmt.Println(fmt.Sprintf("Hashing Power: %d hashes per second", hashPower))
		}
	}
}

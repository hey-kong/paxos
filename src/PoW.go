package main

import (
	"fmt"
	"strings"

	"crypto/sha256"
	"encoding/hex"
)

func main() {
	data := "Hello, world!"
	difficulty := 4
	target := strings.Repeat("0", difficulty) // "0000"

	// 1. Start with nonce at 0
	nonce := 0

	for {
		// 2. Combine the block data and Nonce, i.e., the calculation target is the string
		// "Hello, world!0".
		dataToHash := data + fmt.Sprint(nonce)

		// 3. Apply the SHA-256 hash function to "Hello, world!0" and calculate the hash value.
		hasher := sha256.New()
		hasher.Write([]byte(dataToHash))
		hash := hex.EncodeToString(hasher.Sum(nil))

		// 4. Check if the calculated hash value meets the Difficulty conditions (starting with "0000").
		if strings.HasPrefix(hash, target) {
			// 6. If true, the calculation is complete.
			fmt.Printf("Found a valid hash: %s (Nonce: %d)\n", hash, nonce)
			break
		} else {
			// 5. If not met, increase the Nonce by 1 and repeat steps 2–4.
			nonce++
		}
	}
}

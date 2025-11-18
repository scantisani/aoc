package part1

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Solve() int {
	return MineAdventCoin("iwrupvqb")
}

func MineAdventCoin(prefix string) int {
	suffix := 1

	for {
		hashable := []byte(prefix + strconv.Itoa(suffix))
		possibleHash := md5.Sum(hashable)

		hashAsString := hex.EncodeToString(possibleHash[:])
		if hashAsString[:5] == "00000" {
			return suffix
		}

		suffix++
	}
}

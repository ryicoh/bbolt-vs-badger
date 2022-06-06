package bboltvsbadger

import (
	"crypto/rand"
	"encoding/gob"
	"os"
)

type Pair struct {
	Key   []byte
	Value []byte
}

func genRandByteArray(n int) []byte {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

func newDataset(num int) []Pair {
	dataset := make([]Pair, num)
	for i := 0; i < num; i++ {
		dataset[i] = Pair{
			Key:   genRandByteArray(30),
			Value: genRandByteArray(100),
		}
	}
	return dataset
}

func saveDataset(path string, pairs []Pair) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gob.NewEncoder(file).Encode(pairs)
	if err != nil {
		panic(err)
	}
}

func restoreDataset(path string) []Pair {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var pairs []Pair
	err = gob.NewDecoder(file).Decode(&pairs)
	if err != nil {
		panic(err)
	}

	return pairs
}

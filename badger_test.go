package bboltvsbadger

import (
	"bytes"
	"testing"

	badger "github.com/dgraph-io/badger/v3"
)

var badgerPath = "BenchmarkBadger"
var badgerDataset = "BadgerDataset"
var num = 100

func init() {
	dataset := newDataset(num)
	saveDataset(badgerDataset, dataset)
}

func BenchmarkBadgerPut(b *testing.B) {
	dataset := restoreDataset(badgerDataset)

	opts := badger.DefaultOptions(badgerPath).
		WithSyncWrites(true).
		WithLoggingLevel(badger.ERROR).
		WithMetricsEnabled(false)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < num; i++ {
			err := db.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry(dataset[i].Key, dataset[i].Value)
				return txn.SetEntry(entry)
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkBadgerGet(b *testing.B) {
	dataset := restoreDataset(badgerDataset)

	opts := badger.DefaultOptions(badgerPath).
		WithSyncWrites(true).
		WithLoggingLevel(badger.ERROR).
		WithMetricsEnabled(false)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < num; i++ {
			err := db.View(func(txn *badger.Txn) error {
				item, err := txn.Get(dataset[i].Key)
				if err != nil {
					return err
				}

				return item.Value(func(val []byte) error {
					if !bytes.Equal(val, dataset[i].Value) {
						panic("value does not match")
					}
					return nil
				})
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

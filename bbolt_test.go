package bboltvsbadger

import (
	"bytes"
	"testing"

	"go.etcd.io/bbolt"
)

var bucket = []byte("1")
var bboltPath = "BenchmarkBBolt"
var bboltDataset = "BboltDataset"

func init() {
	dataset := newDataset(num)
	saveDataset(bboltDataset, dataset)
}

func BenchmarkBboltPut(b *testing.B) {
	dataset := restoreDataset(bboltDataset)

	db, err := bbolt.Open(bboltPath, 0644, &bbolt.Options{
		NoSync: false,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < num; i++ {
			err := db.Update(func(tx *bbolt.Tx) error {
				return tx.Bucket(bucket).Put(dataset[i].Key, dataset[i].Value)
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkBboltGet(b *testing.B) {
	dataset := restoreDataset(bboltDataset)

	db, err := bbolt.Open(bboltPath, 0644, &bbolt.Options{
		NoSync: false,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < num; i++ {
			err := db.View(func(tx *bbolt.Tx) error {
				value := tx.Bucket(bucket).Get(dataset[i].Key)
				if !bytes.Equal(value, dataset[i].Value) {
					panic("value does not match")
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

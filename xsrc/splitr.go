package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xbdb"
)

type rec struct {
	K []byte
	V []byte
}
type Rec []rec

func main() {
	f, _ := ioutil.ReadDir("./db")
	for _, fn := range f {
		if fn.Name() != "system.db" {
			fmt.Println(fn.Name())
			//6-merge,8-split
			split(fn.Name(), 8)
		}
	}
}

func split(fn string, ln int) {
	db := xbdb.Opendb("./db/"+fn, 0600)
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			fmt.Println(string(k))
			var rc Rec
			bd := xbdb.Opendb("./db/"+string(k)[0:ln]+".db", 0600)
			b := tx.Bucket(k)
			i := 0
			b.ForEach(func(s, v []byte) error {
				rc = append(rc, rec{s, v})
				i++
				return nil
			})
			Set(k, rc, bd)
			fmt.Println(i, len(rc))
			bd.Close()
		}
		return nil
	})
}

func Set(buck []byte, rc Rec, db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(buck)
		if err != nil {
			return err
		}
		for _, v := range rc {
			err = bucket.Put(v.K, v.V)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

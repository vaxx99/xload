package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xama"
	"github.com/vaxx99/xload/xbdb"
)

func main() {
	bd := xbdb.Opendb("./db/system.db", 0600)
	defer bd.Close()

	w := xama.Redrec{}
	w.Ds = "20170401"
	w.Na = "38272000"
	w.De = "20170410"
	rc := []xama.Block{}
	bb, b := xbdb.Bucket(w, bd)
	fmt.Println(bb)
	os.Exit(0)
	db := xbdb.Opendb("./db/"+b[0][:6]+".db", 0600)
	defer db.Close()
	for i, j := range b {
		fmt.Println(i, j)
		rec := xbdb.Find(j, w, db)
		if len(rec) > 0 {
			rc = append(rc, rec)
		}
	}
	for k, v := range rc {
		fmt.Println(k, v)
	}
}

func Bucket(w xama.Redrec, db *bolt.DB) []string {
	var bc []string
	var bk []string
	var rc []string
	bn := map[string]int{}
	bd := map[string]int{}
	cd := time.Now().Format("20060102")
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("size"))
		if err := b.ForEach(func(k, v []byte) error {
			if string(k) != cd {
				bc = append(bc, string(k))
				bn[string(k)]++
			}
			if len(w.Ds) == 6 && string(k)[:6] == w.Ds {
				//fmt.Println("Ds:", string(k)[:6], string(k))
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.Ds) >= 8 && string(k)[:8] == w.Ds[0:8] {
				//fmt.Println("Ds:", string(k)[:6], string(k))
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.De) == 6 && string(k)[:6] == w.De {
				//fmt.Println("De:", string(k)[:6], string(k))
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.De) >= 8 && string(k)[:8] == w.De[0:8] {
				//fmt.Println("De:", string(k)[:6], string(k))
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.Ds) >= 8 && len(w.De) >= 8 && w.Ds[0:6] == w.De[0:6] {
				ds, e := strconv.Atoi(w.Ds[6:8])
				de, e := strconv.Atoi(w.De[6:8])
				if ds < de && e == nil {
					for i := ds; i <= de; i++ {
						if bd[w.Ds[0:6]+dd(i)] < 1 {
							bk = append(bk, w.Ds[0:6]+dd(i))
						}
						bd[w.Ds[0:6]+dd(i)]++
					}
				}
				//
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	for _, vk := range bk {
		if bn[vk] > 0 {
			rc = append(rc, vk)
		}
	}
	if len(rc) > 0 {
		return rc
	}
	return []string{bc[len(bc)-1]}
}

func dd(d int) string {
	if d < 10 {
		return "0" + strconv.Itoa(d)
	}
	return strconv.Itoa(d)
}

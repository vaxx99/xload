package xbdb

import (
	"hash/fnv"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xama"
)

func Opendb(name string, mod os.FileMode) *bolt.DB {
	db, err := bolt.Open(name, mod, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Set(buck, key, val string, db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(buck))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), []byte(val))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Bet(buck, k string, v []byte, db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(buck))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(k), v)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Rset(recs xama.Block, db *bolt.DB) error {
	e := db.Update(func(tx *bolt.Tx) error {
		for _, j := range recs {

			s := j.Sw + "." + j.Hi + "." + j.Sc + "." +
				j.Na + "." + j.Nb + "." + j.Ds + "." +
				j.De + "." + j.Dr + "." + j.Ot + "." +
				j.It + "." + j.Du + "."

			mn := j.Ds
			if j.Ds == "" {
				mn = j.De
			}
			k := Hash([]byte(s))
			k = mn + ":" + k
			bucket, e := tx.CreateBucketIfNotExists([]byte(mn[0:8]))
			if e != nil {
				return e
			}
			e = bucket.Put([]byte(k), []byte(s))
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	return e
}

func Fget(key string, db *bolt.DB) bool {
	var f bool
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("file"))
		if bucket == nil {
			f = false
			return nil
		}

		val := bucket.Get([]byte(key))
		if val != nil {
			f = true
		} else {
			f = false
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return f
}

func Sget(b, k string, db *bolt.DB) int64 {
	var f int64
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))
		if bucket == nil {
			return nil
		}

		val := bucket.Get([]byte(k))
		if val != nil {
			f, _ = strconv.ParseInt(string(val), 10, 64)
		} else {
			f = 0
		}
		return nil
	})
	return f
}

func Hash(s []byte) string {
	h := fnv.New32a()
	h.Write(s)
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

func Size(cb string, db *bolt.DB) struct{ A, B, C, D int } {
	var s struct{ A, B, C, D int }
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cb))
		b.ForEach(func(k, v []byte) error {
			sc := strings.Split(string(v), ".")
			am := xama.Redrec{Sw: sc[0], Hi: sc[1], Sc: sc[2], Na: sc[3], Nb: sc[4], Ds: sc[5], De: sc[6], Dr: sc[7], Ot: sc[8], It: sc[9], Du: sc[10]}
			if am.Sw == "3810" {
				s.A++
			}
			if am.Sw == "3800" {
				s.B++
			}
			if am.Hi == "si2k" {
				s.C++
			}
			if am.Hi == "es11" {
				s.D++
			}
			return nil
		})
		return nil
	})
	return s
}

//Compare Redrec's
func Comp(s, w xama.Redrec) bool {
	var find []int
	var fc int
	ss := reflect.ValueOf(&s).Elem()
	ws := reflect.ValueOf(&w).Elem()
	wt := ws.Type()
	for i := 0; i < ws.NumField(); i++ {
		sf := ss.Field(i).String()
		df := ws.Field(i).String()
		if df != "" {
			fc += 1
			if strings.Contains(sf, df) && wt.Field(i).Name != "Du" {
				find = append(find, 0)
			}
			if wt.Field(i).Name == "Du" {
				dd, _ := strconv.Atoi(df)
				sd, _ := strconv.Atoi(sf)
				if dd == sd {
					find = append(find, 0)
				}
			}
		}
	}
	var x int
	for i := 0; i < len(find); i++ {
		x += find[i]
	}
	if fc == len(find) {
		return true
	} else {
		return false
	}
}

func Srec(s string) xama.Redrec {
	sc := strings.Split(s, ".")
	return xama.Redrec{Sw: sc[0], Hi: sc[1], Sc: sc[2], Na: sc[3], Nb: sc[4], Ds: sc[5], De: sc[6], Dr: sc[7], Ot: sc[8], It: sc[9], Du: sc[10]}
}

func Dfmt(dt string) string {
	rd := ""
	if len(dt) > 0 {
		rd = dt[6:8] + "." + dt[4:6] + "." + dt[0:4] + " " + dt[8:10] + ":" + dt[10:12] + ":" + dt[12:14]
	}
	return rd
}

func Find(b string, w xama.Redrec, db *bolt.DB) xama.Block {
	var rec xama.Block
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b))
		if err := b.ForEach(func(k, v []byte) error {
			j := Srec(string(v))
			if Comp(j, w) {
				j.Ds = Dfmt(j.Ds)
				j.De = Dfmt(j.De)
				rec = append(rec, j)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return rec
}

func Bucket(w xama.Redrec, db *bolt.DB) (bool, []string) {
	between := false
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
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.Ds) >= 8 && string(k)[:8] == w.Ds[0:8] {
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.De) == 6 && string(k)[:6] == w.De {
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.De) >= 8 && string(k)[:8] == w.De[0:8] {
				if bd[string(k)] < 1 {
					bk = append(bk, string(k))
				}
				bd[string(k)]++
			}
			if len(w.Ds) >= 8 && len(w.De) >= 8 && w.Ds[0:6] == w.De[0:6] && w.Ds[0:8] != w.De[0:8] {
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
				between = true
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
		return between, rc
	}
	return between, []string{bc[len(bc)-1]}
}

func dd(d int) string {
	if d < 10 {
		return "0" + strconv.Itoa(d)
	}
	return strconv.Itoa(d)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/vaxx99/xload/xama"
	"github.com/vaxx99/xload/xbdb"
)

var wd string

func main() {
	t := time.Now()
	fmt.Println("Loader:", t.Format("15:04:05"))
	os.Chdir("./tmp")
	var count int64
	var fnmn string
	buck := map[string]int{}

	fb := xbdb.Opendb("../db/system.db", 0666)
	defer fb.Close()
	f, _ := ioutil.ReadDir(".")
	for _, fn := range f {
		bf, ft := Work(fn.Name())
		tmp := map[string]xama.Block{}
		if !xbdb.Fget(fn.Name(), fb) && bf == true {
			cnt, rp := Recs(fn.Name(), ft)
			for _, j := range rp {
				mn := j.Ds
				if j.Ds == "" {
					mn = j.De
				}
				fnmn = mn[0:8]
				tmp[mn[0:8]] = append(tmp[mn[0:8]], j)
			}
			log.Println("Ok!", ft, fn.Name(), cnt)
			for i, j := range tmp {
				wb := xbdb.Opendb("../db/"+i[0:6]+".db", 0666)
				e := xbdb.Rset(j, wb)
				buck[i[:8]]++
				wb.Close()
				if e != nil {
					log.Println("Alarm!", fn.Name())
					log.Fatal(e)
				}
			}
			count += int64(cnt)
			xbdb.Set("file", fn.Name(), fnmn, fb)
			os.Remove(fn.Name())
		}
	}
	fmt.Printf("%s %.2f\n", "Loader:", time.Now().Sub(t).Seconds())
	st := time.Now()
	fmt.Println("Bucker:", st.Format("15:04:05"))
	for bckn, bcnt := range buck {
		sb := xbdb.Opendb("../db/"+bckn[0:6]+".db", 0666)
		sz := xbdb.Size(bckn, sb)
		v, _ := json.Marshal(sz)
		_ = xbdb.Bet("size", bckn, v, fb)
		sb.Close()
		fmt.Printf("%s %4d %8d %8d %8d %8d\n", bckn, bcnt, sz.A, sz.B, sz.C, sz.D)
	}
	fmt.Printf("%s %.2f\n", "Bucker:", time.Now().Sub(st).Seconds())
}

func Work(fn string) (bool, string) {
	if f, t := xama.Ises(fn); f != false {
		return f, t
	}
	if f, t := xama.Issi(fn); f != false {
		return f, t
	}
	if f, t := xama.Isam(fn); f != false {
		return f, t
	}
	return false, "0"
}

func Recs(fn, ft string) (int, xama.Block) {
	switch ft {
	case "3800":
		_, cnt, rp := xama.Rama(fn)
		return cnt, rp
	case "3810":
		_, cnt, rp := xama.Rama(fn)
		return cnt, rp
	case "si2k":
		cnt, _, _, rp := xama.Si2k(fn)
		return cnt, rp
	case "es11":
		_, cnt, rp := xama.Es11(fn)
		return cnt, rp
	}
	return 0, xama.Block{}
}

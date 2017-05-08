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
    var hv string
    buck := map[string]int{}

    fb := xbdb.Opendb("../db/system.db", 0600)
    defer fb.Close()
    f, _ := ioutil.ReadDir(".")
    for _, fn := range f {
	bf, ft := Work(fn.Name())
	tmp := map[string]xama.Block{}
	hv = xbdb.Fhash(fn.Name())
	if !xbdb.Fget(fn.Name(), hv, fb) && bf == true {
	    cnt, rp := Recs(fn.Name(), ft)
	    for _, j := range rp {
		mn := j.Ds
		if j.Ds == "" {
		    mn = j.De
		}
		tmp[mn[0:8]] = append(tmp[mn[0:8]], j)
	    }
	    log.Println("Ok!", ft, fn.Name(), cnt)
	    for i, j := range tmp {
		wb := xbdb.Opendb("../db/"+i[0:8]+".db", 0600)
		e := xbdb.Rset(j, wb)
		buck[i[:8]]++
		wb.Close()
		if e != nil {
		    log.Println("Alarm!", fn.Name())
		    log.Fatal(e)
		}
	    }
	    count += int64(cnt)
	    xbdb.Set("file", fn.Name(), hv, fb)
	    os.Remove(fn.Name())
	}
    }
    st := time.Now()
    for bckn, bcnt := range buck {
	sb := xbdb.Opendb("../db/"+bckn[0:8]+".db", 0600)
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

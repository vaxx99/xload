package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xbdb"
)

type Base struct{ A, B, C, D, E int }
type Frec struct {
	T  string
	A  int
	VA string
	B  int
	VB string
	C  int
	VC string
	D  int
	VD string
}
type Facs struct {
	DT string
	TM string
	AL Base
	Fc []Frec
}

func dd(d int) string {
	if d < 10 {
		return "0" + strconv.Itoa(d)
	}
	return strconv.Itoa(d)
}

func main() {

	bd := xbdb.Opendb("./db/system.db", 0600)
	defer bd.Close()
	rec := Face(bd)
	buf := new(bytes.Buffer)
	tx := template.New("reps")
	tx, e := template.ParseFiles("./xtmp/reps.tmpl")
	tx.ExecuteTemplate(buf, "reps", rec)
	f, e := os.Create("stat.html")
	defer f.Close()
	f, e = os.OpenFile("stat.html", os.O_WRONLY, 0666)
	defer f.Close()
	if e != nil {
		log.Fatal(e)
	}
	buf.WriteTo(f)
}

func Face(db *bolt.DB) Facs {
	var s struct{ A, B, C, D int }
	var f Facs
	if e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("size"))
		dm := ""
		mv := map[string]int{}
		iv := map[string]int{}
		if e := b.ForEach(func(k, v []byte) error {
			_ = json.Unmarshal(v, &s)
			dt := string(k)
			if dt[:6] == time.Now().Add(-24*time.Hour).Format("200601") {
				mv["nA"] = s.A
				mv["nB"] = s.B
				mv["nC"] = s.C
				mv["nD"] = s.D
				dm = dt[6:8] + "." + dt[4:6] + "." + dt[:4]
				f.Fc = append(f.Fc, Frec{dm, s.A, UpDn(mv["pA"], mv["nA"]),
					s.B, UpDn(mv["pB"], mv["nB"]),
					s.C, UpDn(mv["pC"], mv["nC"]),
					s.D, UpDn(mv["pD"], mv["nD"])})
				mv["pA"] = s.A
				mv["pB"] = s.B
				mv["pC"] = s.C
				mv["pD"] = s.D
				iv["A"] += s.A
				iv["B"] += s.B
				iv["C"] += s.C
				iv["D"] += s.D
			}
			return nil
		}); e != nil {
			return e
		}
		f.DT = dm
		f.TM = time.Now().Format("15:04:05")
		f.AL = Base{iv["A"] + iv["B"] + iv["C"] + iv["D"], iv["A"], iv["B"], iv["C"], iv["D"]}
		return nil
	}); e != nil {
		log.Fatal(e)
	}
	return f
}

func UpDn(vp, vn int) string {

	if vn > vp {
		return "up"
	} else {
		return "dn"
	}
}

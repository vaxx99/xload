package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xama"
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

type Srec struct {
	Rcn int
	Rdr string
	Rec xama.Block
	Old xama.Redrec
}

type Logrec struct {
	A string
	B string
	C int
	D string
	E xama.Redrec
}

func Slog(k string, recs *Logrec, db *bolt.DB) {

	e := db.Update(func(tx *bolt.Tx) error {

		bucket, e := tx.CreateBucketIfNotExists([]byte("logs"))
		if e != nil {
			return e
		}
		v, _ := json.Marshal(recs)
		e = bucket.Put([]byte(k), v)

		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	go Finder()
	for {
		time.Sleep(time.Second * 5)
	}
}

func Finder() {
	log.Println("Start Finder...")
	http.HandleFunc("/", Stat)
	http.HandleFunc("/call", Call)
	http.HandleFunc("/logs", Logs)
	http.Handle("/xcss/", http.StripPrefix("/xcss/", http.FileServer(http.Dir("xcss"))))
	http.Handle("/xtmp/", http.StripPrefix("/xtmp/", http.FileServer(http.Dir("xtmp"))))
	e := http.ListenAndServe(":4444", nil)
	if e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}

func Stat(w http.ResponseWriter, req *http.Request) {
	bd := xbdb.Opendb("./db/system.db", 0600)
	defer bd.Close()
	rec := Face(bd)
	tx := template.New("stat")
	tx, _ = template.ParseFiles("./xtmp/stat.tmpl")
	tx.ExecuteTemplate(w, "stat", rec)
}

func Logs(w http.ResponseWriter, req *http.Request) {
	bd := xbdb.Opendb("./db/system.db", 0600)
	defer bd.Close()
	rec := Rlog(bd)
	tx := template.New("logs")
	tx, _ = template.ParseFiles("./xtmp/logs.tmpl")
	tx.ExecuteTemplate(w, "logs", rec)
}

func logtime() string {
	t := time.Now()
	b := t.UnixNano() % 1e6 / 1e3
	c := strconv.FormatInt(b, 10)
	return t.Format("20060102150405") + c
}

func Call(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		return
	}
	t1 := time.Now()
	st := xama.Redrec{}
	rec := xama.Block{}
	Rc := 0
	Rt := "0.0"
	wt := xama.Redrec{Sw: r.FormValue("sw"), Hi: r.FormValue("hi"), Sc: r.FormValue("sc"), Na: r.FormValue("na"), Nb: r.FormValue("nb"),
		Ds: r.FormValue("ds"), De: r.FormValue("de"), Dr: r.FormValue("dr"), Ot: r.FormValue("ot"), It: r.FormValue("it"), Du: r.FormValue("du")}
	if wt != st {
		bd := xbdb.Opendb("./db/system.db", 0600)
		defer bd.Close()
		bw, bn := xbdb.Bucket(wt, bd)
		db := xbdb.Opendb("./db/"+bn[0][0:6]+".db", 0600)
		defer db.Close()
		for _, nb := range bn {
			wf := wt
			if bw {
				wf.Ds = ""
				wf.De = ""
			}
			rc := xbdb.Find(nb, wf, db)
			if len(rc) > 0 {
				for _, rcc := range rc {
					rec = append(rec, rcc)
				}
			}
			Rc = len(rec)
			Rt = strconv.FormatFloat(time.Now().Sub(t1).Seconds(), 'f', 1, 64)
		}
		Slog(logtime(), &Logrec{time.Now().Format("15:04:05"), r.RemoteAddr, Rc, Rt, wt}, bd)
	}
	if len(rec) > 5000 {
		rec = xama.Block{}
	}
	t := template.New("call")
	t, _ = template.ParseFiles("./xtmp/call.tmpl")
	t.ExecuteTemplate(w, "call", &Srec{Rcn: Rc, Rec: rec, Rdr: Rt, Old: wt})
}

func Face(db *bolt.DB) Facs {
	var s struct{ A, B, C, D int }
	var f Facs
	if e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("size"))
		dm := ""
		dt := ""
		mv := map[string]int{}
		iv := map[string]int{}
		if e := b.ForEach(func(k, v []byte) error {
			_ = json.Unmarshal(v, &s)
			dt = string(k)
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
		f.DT = dt[6:8] + "." + dt[4:6] + "." + dt[:4]
		f.TM = time.Now().Format("15:04:05")
		f.AL = Base{iv["A"] + iv["B"] + iv["C"] + iv["D"], iv["A"], iv["B"], iv["C"], iv["D"]}
		return nil
	}); e != nil {
		log.Fatal(e)
	}
	return f
}

func Rlog(db *bolt.DB) []Logrec {
	var lg Logrec
	var lgs []Logrec
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("logs"))
		b.ForEach(func(k, v []byte) error {
			if strings.Contains(string(k)[:8], time.Now().Format("20060102")) == true {
				e := json.Unmarshal(v, &lg)
				if e != nil {
					log.Printf("Json logs error: %s", e)
				}
				lgs = append(lgs, lg)
			}
			return nil
		})
		return nil
	})
	return lgs
}

func UpDn(vp, vn int) string {

	if vn > vp {
		return "up"
	} else {
		return "dn"
	}
}

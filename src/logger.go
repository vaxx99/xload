package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/vaxx99/xload/xama"
	"github.com/vaxx99/xload/xbdb"
)

type Logrec struct {
	A string
	B string
	C int
	D string
	E xama.Redrec
}
type Reps struct {
	DT string
	RC []Logrec
}

func Rlog(db *bolt.DB) Reps {
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
				lg.B = lg.B[:strings.Index(lg.B, ":")]
				lgs = append(lgs, lg)
			}
			return nil
		})
		return nil
	})
	return Reps{DT: time.Now().Format("02.01.2006"), RC: lgs}
}

func main() {

	bd := xbdb.Opendb("./db/system.db", 0600)
	defer bd.Close()
	rec := Rlog(bd)
	buf := new(bytes.Buffer)
	tx := template.New("logs")
	tx, e := template.ParseFiles("./xtmp/logr.tmpl")
	tx.ExecuteTemplate(buf, "logs", rec)
	f, e := os.Create("logs.html")
	defer f.Close()
	f, e = os.OpenFile("logs.html", os.O_WRONLY, 0666)
	defer f.Close()
	if e != nil {
		log.Fatal(e)
	}
	buf.WriteTo(f)
}

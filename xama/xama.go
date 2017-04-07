package xama

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//Open file
func Open(fn string) (*os.File, error) {
	file, e := os.Open(fn)
	if e != nil {
		log.Println("File open error:", e)
	}
	return file, e
}

//Read file
func Read(file *os.File, bt int) ([]byte, error) {
	data := make([]byte, bt)
	_, e := file.Read(data)
	if e != nil {
		log.Println("File open error:", e)
	}
	return data, e
}

type Redrec struct{ Sw, Hi, Sc, Na, Nb, Ds, De, Dr, Ot, It, Du string }
type Block []Redrec

//5ESS
func A9020(bcd, yy string) string {
	aa := bcd[0:2]
	aa += bcd[2:6]
	aa += bcd[6:12]
	aa += bcd[12:18]
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i : 36]
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i : 70]
	aa += bcd[70:74]
	aa += bcd[74:78]
	aa += bcd[78:80]
	aa += bcd[80:82]
	aa += bcd[82:84]
	aa += bcd[84:86]
	aa += bcd[86:88]
	aa += yy + bcd[89:99]
	aa += bcd[99:100]
	aa += yy + bcd[101:111]
	aa += bcd[111:112]
	aa += bcd[112:114]
	aa += bcd[114:118]
	aa += bcd[118:122]
	aa += bcd[122:126]
	aa += bcd[126:132]
	aa += bcd[132:138]
	aa += bcd[138:140]
	aa += bcd[140:148]
	aa += bcd[148:150]
	aa += bcd[150:152]
	aa += bcd[152:154]
	aa += bcd[154:156]
	aa += bcd[156:158]
	aa += bcd[158:160]
	return aa
}

func A9021(bcd, yy string) string {
	aa := bcd[0:2]                   //Hexadecimal Identifier
	aa += bcd[2:6]                   //Structure Identifier Code
	aa += bcd[6:12]                  //Ticket Number
	aa += bcd[12:18]                 //Sequence Number
	i, _ := strconv.Atoi(bcd[18:20]) //
	aa += bcd[36-i : 36]             //Originating Phone Number
	i, _ = strconv.Atoi(bcd[36:38])  //
	aa += bcd[70-i : 70]             //Terminating Phone Number
	aa += bcd[70:72]                 //Charge Category
	aa += bcd[72:74]                 //Nature of Call
	aa += bcd[74:76]                 //CDA Indicator
	aa += bcd[76:78]                 //LDC Indicator
	aa += bcd[78:80]                 //Service Class of Call
	aa += yy + bcd[81:91]            //Date and Time of Charging Commencement
	aa += bcd[91:92]                 //*ms
	aa += bcd[92:104]                //Date and Time of Call End
	aa += bcd[104:106]               //Cause of Call End
	aa += bcd[106:110]               //Destination
	aa += bcd[110:114]               //Outgoing Trunk Group
	aa += bcd[114:118]               //Incoming Trunk Group
	aa += bcd[118:124]               //Conversation Time
	aa += bcd[124:130]               //Chargeable Duration
	aa += bcd[130:132]               //Class of Rate
	aa += bcd[132:140]               //Fee
	aa += bcd[140:142]               //Trouble Mark
	aa += bcd[142:144]               //Day
	aa += bcd[144:146]               //A-Party Category
	aa += bcd[146:148]               // Type of Call
	aa += bcd[148:150]               //Customer Feature
	aa += bcd[150:152]               //Customer Feature Action
	return aa
}

func A9025(bcd, yy string) string {
	aa := bcd[0:2]   //Hexadecimal Identifier
	aa += bcd[2:6]   //Structure Identifier Code
	aa += bcd[6:12]  //Ticket Number
	aa += bcd[12:18] //Sequence Number
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i : 36] //Originating Phone Number
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i : 70]  //Terminating Phone Number
	aa += bcd[70:72]      //Charge Category
	aa += bcd[72:74]      //Nature of Call
	aa += bcd[74:76]      //CDA Indicator
	aa += bcd[76:78]      //LDC Indicator
	aa += bcd[78:80]      //Service Class of Call
	aa += yy + bcd[81:91] //Date and Time of Charging Commencement
	aa += bcd[91:92]
	aa += yy + bcd[93:103] //Date and Time of Call End
	aa += bcd[103:104]
	aa += bcd[104:106] //Cause of Call End
	aa += bcd[106:110] //Destination
	aa += bcd[110:114] //Outgoing Trunk Group
	aa += bcd[114:118] //Incoming Trunk Group
	aa += bcd[118:124] //Conversation Time
	aa += bcd[124:130] //Chargeable Duration
	aa += bcd[130:132] //Class of Rate
	aa += bcd[132:140] //Fee
	aa += bcd[140:142] //Trouble Mark
	aa += bcd[142:144] //Day
	aa += bcd[144:146] //A-Party Category
	aa += bcd[146:148] //Type of Call
	aa += bcd[148:150] //Bearer Service
	aa += bcd[150:154] //CUG Interlock Code
	aa += bcd[154:156] //COG OA Indicator
	aa += bcd[156:160] //UUI Messages
	aa += bcd[160:162] //Terminating Access
	aa += bcd[162:164] //Network Indicator
	aa += bcd[164:168] //Release Cause
	aa += bcd[168:170] //Supplementary Service Indicator
	return aa
}

func A9026(bcd, yy string) string {
	aa := bcd[0:2]   // Hexadecimal Identifier
	aa += bcd[2:6]   // Structure Identifier Code
	aa += bcd[6:12]  // Ticket Number
	aa += bcd[12:18] // Sequence Number
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i : 36] // Originating Phone Number
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i : 70]   // Terminating Phone Number
	aa += bcd[70:72]       // Charge Category
	aa += bcd[72:74]       // Nature of Call
	aa += bcd[74:76]       // CDA Indicator
	aa += bcd[76:78]       // LDC Indicator
	aa += bcd[78:80]       // Service Class of Call
	aa += yy + bcd[81:91]  // Date and Time of Charging Commencement
	aa += bcd[91:92]       //*ms
	aa += yy + bcd[93:103] // Date and Time of Call End
	aa += bcd[103:104]     //*ms
	aa += bcd[104:106]     // Cause of Call End
	aa += bcd[106:110]     // Destination
	aa += bcd[110:114]     // Outgoing Trunk Group
	aa += bcd[114:118]     // Incoming Trunk Group
	aa += bcd[118:124]     // Conversation Time
	aa += bcd[124:130]     // Chargeable Duration
	aa += bcd[130:132]     // Class of Rate
	aa += bcd[132:140]     // Fee
	aa += bcd[140:142]     // Trouble Mark
	aa += bcd[142:144]     // Day
	aa += bcd[144:146]     // A-Party Category
	aa += bcd[146:148]     // Type of Call
	aa += bcd[148:150]     // Bearer Service
	aa += bcd[150:152]     // Supplementary Service Indicator
	aa += bcd[152:154]     // Supplementary Service Action
	return aa
}

func A9050(bcd string) string {
	aa := bcd[8:10]
	aa += bcd[0:2]
	aa += bcd[2:6]
	aa += bcd[6:10]
	aa += bcd[10:14]
	aa += bcd[14:22]
	aa += bcd[22:28]
	aa += bcd[28:36]
	aa += bcd[36:42]
	aa += bcd[42:46]
	return aa
}

func A9051(bcd string) string {
	aa := bcd[0:2]
	aa += bcd[2:6]
	aa += bcd[6:10]
	aa += bcd[10:14]
	aa += bcd[14:22]
	aa += bcd[22:28]
	aa += bcd[28:36]
	aa += bcd[36:42]
	aa += bcd[42:46]
	aa += bcd[46:54]
	return aa
}

//AA ...
func AA(ad string, yr string) Redrec {
	var a Redrec
	if ad[0:6] == "AA9020" {
		i, e := strconv.Atoi(ad[152:154])
		a.Sw = "3800"
		a.Sc = strconv.Itoa(i)
		a.Hi = ad[2:6]                   //Hexadecimal Identifier
		i, e = strconv.Atoi(ad[18:20])   //Originating Phone Number Digits
		a.Na = ad[36-i : 36]             //Originating Phone Number
		i, e = strconv.Atoi(ad[36:38])   //Terminating Phone Number Digits
		a.Nb = ad[70-i : 70]             //Terminating Phone Number
		a.Ds = yr + ad[89:99]            //Date and Time of Charging Commencement
		a.De = yr + ad[101:111]          //Date and Time of Call End
		a.Dr = ad[114:118]               //Destination
		a.Ot = ad[118:122]               //Outgoing Trunk Group
		a.It = ad[122:126]               //Incoming Trunk Group
		i, e = strconv.Atoi(ad[126:130]) //Conversation Time
		j, e := strconv.Atoi(ad[130:132])
		a.Du = strconv.Itoa((i * 60) + j)
		Err(e)
	}
	if ad[0:6] == "AA9021" {
		i, e := strconv.Atoi(ad[144:146])
		a.Sw = "3800"
		a.Sc = strconv.Itoa(i)
		a.Hi = ad[2:6]                 //Hexadecimal Identifier
		i, e = strconv.Atoi(ad[18:20]) //Originating Phone Number Digits
		a.Na = ad[36-i : 36]           //Originating Phone Number
		i, e = strconv.Atoi(ad[36:38]) //Terminating Phone Number Digits
		a.Nb = ad[70-i : 70]           //Terminating Phone Number
		a.Ds = yr + ad[81:91]          //Date and Time of Charging Commencement
		//a.De = yr + ad[92:104]        //Date and Time of Call End
		a.Dr = ad[106:110]               //Destination
		a.Ot = ad[110:114]               //Outgoing Trunk Group
		a.It = ad[114:118]               //Incoming Trunk Group
		i, e = strconv.Atoi(ad[124:130]) //Chargeable Duration
		a.Du = strconv.Itoa(i)

		Err(e)
	}

	if ad[0:6] == "AA9025" {
		i, e := strconv.Atoi(ad[144:146])
		a.Sw = "3800"
		a.Sc = strconv.Itoa(i)
		a.Hi = ad[2:6]                 //Hexadecimal Identifier
		i, e = strconv.Atoi(ad[18:20]) //Originating Phone Number Digits
		a.Na = ad[36-i : 36]           //Originating Phone Number
		i, e = strconv.Atoi(ad[36:38]) //Terminating Phone Number Digits
		a.Nb = ad[70-i : 70]           //Terminating Phone Number
		a.Ds = yr + ad[81:91]
		a.De = yr + ad[93:103]
		a.Dr = ad[106:110]               //Destination
		a.Ot = ad[110:114]               //Outgoing Trunk Group
		a.It = ad[114:118]               //Incoming Trunk Group
		i, e = strconv.Atoi(ad[118:122]) //Conversation Time
		j, e := strconv.Atoi(ad[122:124])
		a.Du = strconv.Itoa((i * 60) + j)
		Err(e)
	}
	if ad[0:6] == "AA9026" {
		i, e := strconv.Atoi(ad[144:146])
		a.Sw = "3800"
		a.Sc = strconv.Itoa(i)
		a.Hi = ad[2:6]                 //Hexadecimal Identifier
		i, e = strconv.Atoi(ad[18:20]) //Originating Phone Number Digits
		a.Na = ad[36-i : 36]           //Originating Phone Number
		i, e = strconv.Atoi(ad[36:38]) //Terminating Phone Number Digits
		a.Nb = ad[70-i : 70]           //Terminating Phone Number
		a.Ds = yr + ad[81:91]
		//a.De = yr + ad[93:103]
		a.Dr = ad[106:110]               //Destination
		a.Ot = ad[110:114]               //Outgoing Trunk Group
		a.It = ad[114:118]               //Incoming Trunk Group
		i, e = strconv.Atoi(ad[124:130]) //Chargeable Duration
		a.Du = strconv.Itoa(i)
		Err(e)
	}
	if ad[0:6] == "AA0003" {
		a.Sw = "3810"
		a.Hi = ad[2:6]            //Hexadecimal Identifier
		a.Ot = ad[6:10]           //Outgoing Trunk Group
		a.It = ad[10:14]          //Incoming Trunk Group
		a.Ds = Dtt(yr, ad[15:25]) //Date and Time of Charging Commencement
		a.Dr = ad[28:32]          //Destination
		i, e := strconv.Atoi(ad[32:34])
		a.Nb = ad[66-i : 66]      //Terminating Phone Number
		a.De = Dtt(yr, ad[71:81]) //Date and Time of Call End
		a.Du = Diff(a.Ds, a.De)   //Duration
		Err(e)
	}
	return a
}

//Rama -
func Rama(fn string) (string, int, Block) {
	var recs Block
	ft, mT, rc, j, file, e := Finfo(fn)
	yy := mT[0:4]
	Err(e)
	defer file.Close()
	fe := false
	tf := ""
	i := 0

	if ft == "AMA" && fe != true {
		tf = "3800"
		for {
			data, e := Read(file, j)
			j = Next(data[len(data)-3])
			Err(e)
			ad := H2bcd(data)

			if ad[0:6] == "AA9020" {
				b := AA(ad, yy)
				recs = append(recs, b)
			}
			if ad[0:6] == "AA9021" {
				b := AA(ad, yy)
				recs = append(recs, b)
			}
			if ad[0:6] == "AA9025" {
				b := AA(ad, yy)
				recs = append(recs, b)
			}
			if ad[0:6] == "AA9026" {
				b := AA(ad, yy)
				recs = append(recs, b)
			}
			if j == 0 {
				//exit on "EOF"
				return tf, rc, recs
			}
			i++
		}
	}

	if ft == "IAD" && fe != true {
		tf = "3810"
		i := 0
		for i < rc {
			data, e := Read(file, 42)
			Err(e)
			ad := H2bcd(data)
			if ad[0:6] == "AA0003" {
				b := AA(ad, yy)
				recs = append(recs, b)
			} else {
				_, e = Read(file, 32)
				Err(e)
			}
			i++
		}
	}
	return tf, rc, recs
}

//ABor - AMA begin of record
func ABor(dt string) string {
	return "201" + dt[31:36] + dt[37:43]
}

//AEor - AMA end of record
func AEor(dt string) int {
	rc, _ := strconv.Atoi(dt[54:62])
	return rc
}

//IBor - IAD begin of record
func IBor(dt string) string {
	return "20" + dt[8:20]
}

//IEor - IAD end of record
func IEor(dt string) int {
	rc, _ := strconv.Atoi(dt[20:28])
	return rc
}

func Diff(ds, de string) string {
	// setup a start and end time
	if ds == "" {
		return ""
	}
	if de == "" {
		return ""
	}
	createdAt, _ := time.Parse("20060102150405", ds)
	expiresAt, _ := time.Parse("20060102150405", de)
	// get the diff
	diff := expiresAt.Sub(createdAt).Nanoseconds() / 1000000000
	diff = int64(diff)
	dif := strconv.FormatInt(diff, 10)
	return dif
}

func Dtt(yr, ds string) string {
	// setup a start and end time
	d := yr + ds
	dss, _ := strconv.Atoi(ds)

	if dss == 0 {
		d = ""
	}
	return d

}

//Err - error
func Err(e error) {
	if e != nil {
		panic(e)
	}
}

//H2i - hex to int
func H2i(dt string) int {
	ft, _ := strconv.Atoi(dt)
	return ft
}

func H2int(hexStr string) int {
	// base 16 for hexadecimal
	res, _ := strconv.ParseInt(hexStr, 16, 64)
	return int(res)
}

//H2bcd - hex to bcd
func H2bcd(dt []byte) string {
	hd := strings.ToUpper(hex.EncodeToString(dt))
	return hd
}

//Next record size
func Next(dt byte) int {
	// base 16 for hexadecimal
	res := int(dt)
	return res
}

//Finfo BCD
func Finfo(fn string) (string, string, int, int, *os.File, error) {
	var j int
	var mT string
	ft := "NIL"
	rc := 0
	file, e := os.Open(fn)
	Err(e)
	st, e := file.Stat()
	fs := st.Size()
	Err(e)
	dt := make([]byte, 4)
	_, e = file.Read(dt)
	Err(e)
	hd := strings.ToUpper(hex.EncodeToString(dt))
	if strings.Contains(hd, "001B0000") == true {
		ft = "AMA"
		_, e = file.Seek(0, 0)
		data, e := Read(file, 31)
		Err(e)
		dt := H2bcd(data)
		mT = ABor(dt)
		j = Next(data[len(data)-3])
		_, e = file.Seek(st.Size()-31, 0)
		Err(e)
		data, e = Read(file, 31)
		Err(e)
		dt = H2bcd(data)
		rc = AEor(dt)
		_, e = file.Seek(31, 0)
		Err(e)
	}
	if strings.Contains(hd, "01000000") == true {
		ft = "IAD"
		_, e = file.Seek(0, 0)
		data, e := Read(file, 2048)
		Err(e)
		dt := H2bcd(data)
		mT = IBor(dt)
		_, e = file.Seek(fs-2048, 0)
		Err(e)
		data, e = Read(file, 2048)
		Err(e)
		dt = H2bcd(data)
		rc = IEor(dt)
		_, e = file.Seek(2048, 0)
		Err(e)
	}
	return ft, mT, rc, j, file, e
}

func Isam(fn string) (bool, string) {
	if info, err := os.Stat(fn); err == nil && info.IsDir() {
		return false, "0"
	}
	f, _ := os.Open(fn)
	defer f.Close()
	data, _ := Read(f, 31)
	ad := H2bcd(data)
	if ad[8:14] == "AA9050" {
		return true, "3800"
	}
	if ad[0:2] == "01" {
		return true, "3810"
	}
	return false, "0"
}

//SI2K
type SFlag struct{ f01, f02, f03, f04, f05, f06, f07, f08, f09, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19 int64 }
type Frec struct {
	IDR  int64
	IDC  int64
	FLG  SFlag
	SLS  int64
	CHS  int64
	ZCL  int64
	SLL  int64
	ZCD  string
	SNC  string
	P100 P100
	P102 P102
	P103 P103
	P104 P104
	P105 P105
	P106 P106
	P107 P107
	P108 P108
	P109 P109
	P110 P110
	P111 P111
	P112 P112
	P113 P113
	P114 P114
	P115 P115
	P116 P116
	P119 P119
	P121 P121
}

type P100 struct {
	IDI int64
	CNL int64
	CNC string
}

type P102 struct {
	IDI int64
	DTS string
	F1  int64
}

type P103 struct {
	IDI int64
	DTE string
	F1  int64
}

type P104 struct {
	IDI int64
	CNT int64
}

type P105 struct {
	IDI int64
	SVC int64
	TVC int64
}

type P106 struct {
	IDI int64
	SVC int64
}

type P107 struct {
	IDI int64
	SVC int64
}

type P108 struct {
	IDI int64
	TPE int64
	SVC int64
}

type P109 struct {
	IDI int64
	CNL int64
	CNC string
}

type P110 struct {
	IDI int64
	CAT int64
}

type P111 struct {
	IDI int64
	DIR int64
}

type P112 struct {
	IDI int64
	CFC int64
}

type P113 struct {
	IDI int64
	TGN int64
	SLN int64
	MDN int64
	PTN int64
	CHN int64
}

type P114 struct {
	IDI int64
	TGN int64
	SLN int64
	MDN int64
	PTN int64
	CHN int64
}

type P115 struct {
	IDI int64
	DUR int64
}

type P116 struct {
	IDI int64
	BTL int64
	CRC int64
}

type P119 struct {
	IDI int64
	BTL int64
	CNL int64
	CNC string
}

type P121 struct {
	IDI int64
	BTL int64
	COI int64
	CNC string
}

func Flags(s string) SFlag {
	var f []int64
	for _, j := range s {
		res, _ := strconv.ParseInt(string(j), 10, 64)
		f = append(f, res)
	}
	return SFlag{f[7], f[6], f[5], f[4], f[3], f[2], f[1], f[7], f[15], f[14], f[13], f[12], f[11], f[10], f[9], f[8], f[18], f[17], f[16]}
}

func S2002rec(sw string, srec Frec) Redrec {
	var rec Redrec
	rec.Sw = sw
	rec.Hi = "si2k"
	rec.Sc = strconv.Itoa(int(srec.P110.CAT))
	rec.Na = srec.SNC
	rec.Nb = srec.P100.CNC
	rec.Ds = srec.P102.DTS
	rec.De = srec.P103.DTE
	rec.Dr = strconv.Itoa(int(srec.P111.DIR))
	rec.It = strconv.Itoa(int(srec.P113.TGN))
	rec.Ot = strconv.Itoa(int(srec.P114.TGN))
	rec.Du = strconv.FormatFloat(float64(srec.P115.DUR)/1000, 'f', 0, 64)
	return rec
}

func Si2k(fn string) (int, string, string, Block) {
	f, _ := os.Open(fn)
	var cnt int
	var sw, mtm string
	var rec Frec
	var Rec Block
	defer f.Close()

	for {
		head := make([]byte, 3)
		_, e := f.Read(head)

		if e != nil {
			break
		}

		i := b2i(head[1:3])
		switch head[0] {
		case 200:
			b := make([]byte, i-3)
			_, e = f.Read(b)
			rec = S200(b, i-3)
			cnt++
			sw = fn[1:5]
			Rec = append(Rec, S2002rec(sw, rec))
			if len(rec.P102.DTS) > 0 {
				mtm = rec.P102.DTS
			}
		case 210:
			b := make([]byte, 13)
			_, e = f.Read(b)

		case 211:
			b := make([]byte, 13)
			_, e = f.Read(b)

		case 212:
			b := make([]byte, 6)
			_, e = f.Read(b)
		}
	}
	return cnt, sw, mtm, Rec
}

func S200(b []byte, bs int64) Frec {
	var srec Frec
	//Индекс записи
	srec.IDR = b2i(b[0:4])
	//Идентификатор вызова
	srec.IDC = b2i(b[4:8])
	//Flags
	fb := Oct(b[8]) + Oct(b[9]) + Oct(b[10])
	srec.FLG = Flags(fb)

	bc := Oct(b[11])
	//Последовательность
	a, _ := strconv.ParseInt(bc[:4], 2, 8)
	//Состояние учета	стоимости
	c, _ := strconv.ParseInt(bc[4:], 2, 8)
	srec.SLS = a
	srec.CHS = c

	bc = Oct(b[12])
	//Длина кода зоны
	d, _ := strconv.ParseInt(bc[:3], 2, 8)
	//Длина списочного номера
	f, _ := strconv.ParseInt(bc[3:], 2, 8)
	srec.ZCL = d
	srec.SLL = f
	//Bytes count
	btz := Bts(d)
	btn := Bts(f)
	//Код зоны
	srec.ZCD = H2c(b[13 : 13+btz])[0:d]
	//Списочный номер абонента
	srec.SNC = H2c(b[13 : 13+btz+btn])[0 : d+f]

	//dynamic part start byte
	nb := 13 + int((float64(d)+float64(f))/2+0.5)
	for nb < int(bs) {
		id, _ := strconv.ParseInt(Oct(b[nb]), 2, 64)
		nb = Dynp(id, nb, b, &srec)
		//fmt.Println(id, nb, bs, &srec)
	}

	return srec
}

func b2i(b []byte) int64 {
	hd := strings.ToUpper(hex.EncodeToString(b))
	res, _ := strconv.ParseInt(hd, 16, 64)
	return res
}

func Oct(b byte) string {
	return fmt.Sprintf("%08b", b)
}

func Bts(d int64) int {
	return int(float64(d)/2 + 0.5)
}

func H2c(dt []byte) string {
	hd := strings.ToUpper(hex.EncodeToString(dt))
	return hd
}

func Dynp(id int64, nb int, b []byte, rec *Frec) int {
	switch id {
	case 100:
		var st P100
		st.IDI = id
		dc, _ := strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		st.CNL = dc
		btb := Bts(dc)
		st.CNC = H2c(b[nb+2 : nb+2+btb])[0:int(st.CNL)]
		nb = nb + 2 + btb
		rec.P100 = st
		return nb
	case 102:
		var st P102
		st.IDI = id
		bc := Oct(b[nb+9])
		f1, _ := strconv.ParseInt(bc[7:], 2, 64)
		st.F1 = f1
		st.DTS = dates(b[nb+1 : nb+8])
		nb = nb + 9
		rec.P102 = st
		return nb
	case 103:
		var st P103
		st.IDI = id
		bc := Oct(b[nb+9])
		f1, _ := strconv.ParseInt(bc[7:], 2, 64)
		st.F1 = f1
		st.DTE = dates(b[nb+1 : nb+8])
		nb = nb + 9
		rec.P103 = st
		return nb
	case 104:
		var st P104
		bc := Oct(b[nb+1]) + Oct(b[nb+2]) + Oct(b[nb+3])
		//Идентификатор информационного элемента (104)
		st.IDI = id
		cnt, _ := strconv.ParseInt(bc, 2, 64)
		//Количество тарифных импульсов
		st.CNT = cnt
		nb = nb + 4
		rec.P104 = st
		return nb
	case 105:
		var st P105
		//Идентификатор информационного элемента (105)
		st.IDI = id
		//Базовая услуга
		st.SVC, _ = strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		//Телеслужбы
		st.TVC, _ = strconv.ParseInt(Oct(b[nb+2]), 2, 64)
		nb = nb + 3
		rec.P105 = st
		return nb
	case 106:
		var st P106
		//Идентификатор информационного элемента (106)
		st.IDI = id
		//Дополнительная услуга
		st.SVC, _ = strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		nb = nb + 2
		rec.P106 = st
		return nb
	case 107:
		var st P107
		st.IDI = id
		st.SVC, _ = strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		nb = nb + 2
		rec.P107 = st
		return nb
	case 108:
		var st P108
		st.IDI = id
		st.TPE = b2i(b[nb+1 : nb+2])
		st.SVC = b2i(b[nb+2 : nb+3])
		nb = nb + 3
		rec.P108 = st
		return nb
	case 109:
		var st P109
		st.IDI = id
		st.CNL = b2i(b[nb+1 : nb+2])
		btb := Bts(st.CNL)
		st.CNC = H2c(b[nb+2 : nb+2+btb])[0:int(st.CNL)]
		nb = nb + 2 + btb
		rec.P109 = st
		return nb
	case 110:
		var st P110
		//Идентификатор информационного элемента (110)
		st.IDI = id
		//Исходящая категория
		st.CAT, _ = strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		nb = nb + 2
		rec.P110 = st
		return nb
	case 111:
		var st P111
		st.IDI = id
		//Тарифное направление
		st.DIR, _ = strconv.ParseInt(Oct(b[nb+1]), 2, 64)
		nb = nb + 2
		rec.P111 = st
		return nb
	case 112:
		var st P112
		st.IDI = id
		st.CFC = b2i(b[nb+1 : nb+3])
		nb = nb + 2
		rec.P112 = st
		return nb
	case 113:
		var st P113
		st.IDI = id
		st.TGN = b2i(b[nb+1 : nb+3])
		st.SLN = b2i(b[nb+3 : nb+5])
		st.MDN = b2i(b[nb+5 : nb+6])
		st.PTN = b2i(b[nb+6 : nb+8])
		st.CHN = b2i(b[nb+8 : nb+9])
		rec.P113 = st
		nb = nb + 9
		return nb
	case 114:
		var st P114
		st.IDI = id
		st.TGN = b2i(b[nb+1 : nb+3])
		st.SLN = b2i(b[nb+3 : nb+5])
		st.MDN = b2i(b[nb+5 : nb+6])
		st.PTN = b2i(b[nb+6 : nb+8])
		st.CHN = b2i(b[nb+8 : nb+9])
		rec.P114 = st
		nb = nb + 9
		return nb
	case 115:
		var st P115
		st.IDI = id
		st.DUR = b2i(b[nb+1 : nb+5])
		rec.P115 = st
		nb = nb + 5
		return nb
	case 116:
		var st P116
		st.IDI = id
		st.BTL = b2i(b[nb+1 : nb+2])
		st.CRC = b2i(b[nb+2 : nb+4])
		rec.P116 = st
		nb = nb + 4
		return nb
	case 119:
		var st P119
		st.IDI = id
		st.BTL = b2i(b[nb+1 : nb+2])
		st.CNL = b2i(b[nb+2 : nb+3])
		btb := Bts(st.CNL)
		st.CNC = H2c(b[nb+3 : nb+3+btb])[0:int(st.CNL)]
		nb = nb + 3 + btb
		rec.P119 = st
		return nb
	case 121:
		var st P121
		st.IDI = id
		st.BTL = b2i(b[nb+1 : nb+2])
		st.COI = b2i(b[nb+2 : nb+4])
		st.CNC = Oct(b[nb+5])
		nb = nb + 5
		rec.P121 = st
		return nb
	}
	fmt.Println("IDDD:", id)
	os.Exit(0)
	return 0
}

func dates(b []byte) string {
	rd := ""
	if len(b) > 0 {
		rd = "20" + dd(int(b[0])) + dd(int(b[1])) + dd(int(b[2])) + dd(int(b[3])) +
			dd(int(b[4])) + dd(int(b[5])) //+ dd(int(b[6]))
	}
	return rd
}

func dd(d int) string {
	if d < 10 {
		return "0" + strconv.Itoa(d)
	}
	return strconv.Itoa(d)
}

func Issi(fn string) (bool, string) {
	f, _ := os.Open(fn)
	defer f.Close()
	s, _ := f.Stat()

	if s.Size() < 20 {
		return false, "0"
	}
	data, _ := Read(f, 1)
	ad := H2c(data)
	switch ad {
	case "C8":
		return true, "si2k"
	case "D2":
		return true, "si2k"
	case "D3":
		return true, "si2k"
	case "D4":
		return true, "si2k"
	}

	return false, "0"
}

//ES11
type Fields struct {
	FN string
	FT string
	FS int
}

type Esrec struct {
	RECNUMB    string
	STNAME     string
	LINETYPE   string
	LINECODE   string
	AREAOFFSET string
	LINECODETO string
	OLDSTATUS  string
	NEWSTATUS  string
	TALKFLAGS  string
	CAUSE      string
	ISUPCAT    string
	ENDDATE    string
	ENDTIME    string
	DURATION   string
	SUBSTO     string
	SUBSFROM   string
	REDIRSUBS  string
	CONNSUBS   string
	TALKCOMM   string
}

func Head(f *os.File) (string, uint32, uint16, uint16, []Fields, error) {
	v, e := Read(f, 1)
	//Date modified 3
	v, e = Read(f, 3)
	//YY
	yy := strconv.Itoa(1900 + int(v[0]))
	mm := dd(int(v[1]))
	dd := dd(int(v[2]))
	dt := dd + "." + mm + "." + yy
	//Number of records 4-7 (4)
	v, e = Read(f, 4)
	rn := binary.LittleEndian.Uint32(v)
	//Number of bytes in header 8-9 (2)
	v, e = Read(f, 2)
	hb := binary.LittleEndian.Uint16(v)
	//Number of bytes in the record 10-11 (2)
	v, e = Read(f, 2)
	rb := binary.LittleEndian.Uint16(v)
	//12-14 	3 bytes 	Reserved bytes.
	v, e = Read(f, 3)
	//15-27 	13 bytes 	Reserved for dBASE III PLUS on a LAN.
	v, e = Read(f, 13)
	//28-31 	4 bytes 	Reserved bytes.
	v, e = Read(f, 4)
	//32-n 	32 bytes 	Field descriptor array (the structure of this array is each shown below)
	var br int
	var fld []Fields

	for br != 13 {
		v, e = Read(f, 32)
		br = int(v[0])
		if br != 13 {
			fld = append(fld, Fields{string(v[0:11]), string(v[11:12]), int(v[16])})
		}
	}
	f.Seek(0, 0)
	return dt, rn, hb, rb, fld, e
}

func es2rec(erec *Esrec) Redrec {
	var rec Redrec
	ds, de, dr := esdates(erec)
	rec.Sw = Sw(erec)
	rec.Hi = "es11"
	rec.Sc = Bcat(erec.ISUPCAT)
	rec.Na = erec.SUBSFROM
	rec.Nb = erec.SUBSTO
	rec.Ds = ds
	rec.De = de
	rec.Dr = erec.OLDSTATUS + erec.NEWSTATUS
	rec.It = erec.LINECODE
	rec.Ot = erec.LINECODETO
	rec.Du = dr
	return rec
}

func Sw(rec *Esrec) string {
	switch rec.STNAME[0:3] {
	case "023":
		return "3846"
	case "024":
		return "3847"
	case "025":
		return "3856"
	case "026":
		return "3853"
	case "027":
		return "3857"
	case "028":
		return "3852"
	case "076":
		return "3859"
	case "088":
		return "3844"
	case "101":
		return "3850"
	case "102":
		return "3841"
	case "431":
		return "3855"
	}
	return "38xx"
}

func Bcat(sc string) string {
	cat, _ := strconv.ParseInt(sc, 16, 64)
	switch cat {
	case 10:
		return "1"
	case 11:
		return "4"
	case 12:
		return "8"
	case 13:
		return "TC"
	case 15:
		return "9"
	case 224:
		return "0"
	case 225:
		return "2"
	case 226:
		return "5"
	case 227:
		return "7"
	case 228:
		return "3"
	case 229:
		return "6"
	}
	return "XX"
}

func Ises(fn string) (bool, string) {
	f, _ := os.Open(fn)
	v, _ := Read(f, 2)
	a := int(v[0])
	b := 1900 + int(v[1])
	defer f.Close()
	if a == 3 && b == time.Now().Add(-24*time.Hour).Year() {
		return true, "es11"
	}
	return false, "0"
}

func Es11(fn string) (string, int, Block) {
	var Rec Block
	f, e := Open(fn)
	defer f.Close()
	dt, rn, hb, rb, fd, e := Head(f)

	if e != nil {
		panic(e)
	}
	f, e = Open(fn)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	v, _ := Read(f, int(hb))
	var rec Esrec

	recVal := reflect.ValueOf(&rec).Elem()
	for i := 0; i < int(rn); i++ {
		v, e = Read(f, int(rb))
		if e != nil {
			panic(e)
		}
		sb := 1
		eb := 1
		for a, b := range fd {
			eb = sb + b.FS
			s := strings.Replace(string(v[sb:eb]), " ", "", -1)
			recVal.Field(a).SetString(s)
			sb = eb
		}
		j := es2rec(&rec)
		Rec = append(Rec, j)
	}
	return dt, int(rn), Rec
}

func esdates(rec *Esrec) (string, string, string) {
	dt := rec.ENDDATE
	tm := rec.ENDTIME
	de := dt + " " + tm
	dr := s2i(rec.DURATION)
	te, _ := time.Parse("20060102 15:04:05", de)
	ts := te.Add(time.Second * time.Duration(dr) * -1)
	de = te.Format("20060102150405")
	ds := ts.Format("20060102150405")
	return ds, de, strconv.Itoa(dr)
}

func s2i(s string) int {
	s = strings.Replace(s, " ", "", -1)
	a, _ := strconv.Atoi(s)
	return a
}

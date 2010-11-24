package main

//import "fmt"
import "time"
import "os"
import "bufio"
import "utf8"
import "unicode"
//import "strings"
import "strconv"

type t_account struct {
    name string
    fullName string
}

type t_ledger_item struct {
    account *t_account
    summa int64
}

type t_ledger_tran struct {
    date time.Time
    description string
    items []*t_ledger_item
}

type t_ledger_file struct {
    Tran []*t_ledger_tran
    Accounts []*t_account
    curT *t_ledger_tran
    curY int64
    curM int
}



func DecodeDate(s string, dt* time.Time) int { // cлайс на входе
    S := utf8.NewString(s)
    var err os.Error

    if S.RuneCount()==2 {   // тока день
        dt.Day, err = strconv.Atoi(S.Slice(0,2))
            if err==os.ERANGE {return(-1)}
        return(0)
    }
    if S.RuneCount()==5 {   // тока месяц день
        dt.Day, err = strconv.Atoi(S.Slice(3,5))
            if err==os.ERANGE {return(-1)}
        dt.Month, err = strconv.Atoi(S.Slice(0,2))
            if err==os.ERANGE {return(-1)}
        return(0)
    }
    if S.RuneCount()==10 {   // все
        dt.Year, err = strconv.Atoi64(S.Slice(0,4))
            if err==os.ERANGE {return(-1)}
        dt.Month, err = strconv.Atoi(S.Slice(5,7))
            if err==os.ERANGE {return(-1)}
        dt.Day, err = strconv.Atoi(S.Slice(8,10))
            if err==os.ERANGE {return(-1)}
        return(0)
    }
    return(-10)
}

func (lf *t_ledger_file) ParseTran(s string) *t_ledger_tran{
    lf.curT = new(t_ledger_tran)
    lf.curT.items = make([]*t_ledger_item,0)
    lf.curT.date.Year = lf.curY
//    lf.curT.date =  new(time.Time)
    S := utf8.NewString(s)
    var i int
    for i=0; i<S.RuneCount(); i++{
        if unicode.IsSpace(S.At(i)){ // конец даты
            DecodeDate(S.Slice(0,i), &lf.curT.date)
            break
            }
    }
    lf.curT.description = S.Slice(i,S.RuneCount())
    print("add tran: ", lf.curT.date.String() ,lf.curT.description)
    lf.Tran = append(lf.Tran,lf.curT)
    return(lf.curT)
}
func (lf *t_ledger_file) ParseTranItem(s string) *t_ledger_item{
    item := new(t_ledger_item)
    lf.curT.items = append(lf.curT.items,item)
    return(item)
}


func ledgerParseFile(fileName string) *t_ledger_file {
    f, err := os.Open(fileName, os.O_RDONLY, 0)
    if err != nil {
      return nil
    }
    defer f.Close()
    lf := new(t_ledger_file)
    lf.Tran = make([]*t_ledger_tran,0)
    lf.Accounts = make([]*t_account,0)
    var s string
    //var n int = 1
    fr := bufio.NewReader(f)
    s, err = fr.ReadString('\n')
    for s!="" {
        S := utf8.NewString(s)

        if S.At(0)==int(';') {
            println("// ",s)
        }
        if S.At(0)==int('Y') { // текущий год
            lf.curY, err = strconv.Atoi64(S.Slice(1,5))
        }

        if unicode.IsDigit(S.At(0)) { // tran begin
            lf.ParseTran(s)
        }
        if unicode.IsSpace(S.At(0)) { // tran item
            lf.ParseTranItem(s)
        }
        s, err = fr.ReadString('\n')
    }

    println(lf.Tran)
    return lf
}

func main() {
//        n, err := fmt.Println("hello, world!")
//        fmt.Println(n, err)
        ledgerParseFile("test.txt")
        println("End")
}

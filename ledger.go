package main

//import "fmt"
import "time"
import "os"
import "bufio"
import "utf8"
import "unicode"
//import "strings"
import "strconv"
import "regexp"

type t_account struct {
    name string // как в файле Траты:Вовка:Стрижка
    lastname string // Стрижка
    fullName string //Траты:Вовка:Стрижка:Валюта
    valuta string // валюта
    parent *t_account // выше по уровню
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
    Accounts map[string] *t_account
    curT *t_ledger_tran
    curY int64
    curM int
}

func (lf *t_ledger_file) findOrCreateAccount(s string) *t_account{
    var acc *t_account
    acc = lf.Accounts[s]
    if acc!=nil {// нашли
        return acc
    }
    r,err := regexp.Compile(":")
    a := r.FindStringSubmatch(s)
    println(a,err)
    // создаем
    acc = new(t_account)
    return acc
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
    lf.curT.date.Month = lf.curM
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

func (lf *t_ledger_file) ParseAccount(s string) *t_account{
    var acc *t_account // вычислить счет найти если нет создать
    return (acc)
}

func ParseSumma(s string) int64{
    S := utf8.NewString(s)
    var i int
    for i=0; i<S.RuneCount(); i++{
        if unicode.IsSpace(S.At(i)){ // конец даты
            //DecodeDate(S.Slice(0,i), &lf.curT.date)
            break
            }
    }
    return(0)
}

func (lf *t_ledger_file) ParseTranItem(s string) *t_ledger_item{
    item := new(t_ledger_item)
    lf.curT.items = append(lf.curT.items,item)
    item.account = lf.ParseAccount(s) // чистый аккаунт сюда
    item.summa = ParseSumma(s) // чистая сумма
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
    lf.Accounts = make(map[string]*t_account,0)
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
        if S.At(0)==int('M') { // текущий месяц
            lf.curM, err = strconv.Atoi(S.Slice(1,3))
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
        ledgerParseFile("test.txt")
        println("End")
        s := "Ghbdtn dctv: n bgf n :juJприВет:34261"
//        r,err := regexp.Compile(":([123456789]+)")
        r,err := regexp.Compile(":([a-zA-Z0-9а-яА-Я]+)")
        a := r.FindString(s)
        println(a,err)
}

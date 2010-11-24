package main
import "fmt"
import "strconv"
import "os"
import "bufio"
import "log"
import "big"

type s_a struct {
    a int
    b string
  }

type i_test interface {
    test() int
}

func (s *s_a) test() int{
    fmt.Print(s.a);
    s.b = "666"
    return(0);
}

func iTest(i i_test) int {
    return i.test()
}

func main() {
  ss := new(s_a)
  ss.a = 24
  ss.b = "45"
  ss.b = ss.b+fmt.Sprint(ss.a);
  ss.test();
  n, err := fmt.Print("hello");
  n = n + 1;
  fmt.Print(err);
  ss.a, err = strconv.Atoi(ss.b);
  fmt.Print(ss)
  iTest(ss);

  f, err := os.Open("c:\\my.txt", os.O_RDONLY, 0)
  if err != nil {
      return
  }
  defer f.Close()
  var s string
//  n1, err := fmt.Fscanln(f,"%s",s)
  fr := bufio.NewReader(f)
  s, err = fr.ReadString('\n')
  fmt.Println(s)
  log.Println("log message");

  summa := big.NewInt(0)
  summa.SetString("123454567891234545678912345456789",0)
  summa2 := big.NewInt(0)
  summa2.SetString("7",0)
  summa3 := big.NewInt(0)
  summa3.Exp(summa, summa2,nil)
  fmt.Println(summa3)
  _,_ = 4*5, "test"
  _ = "Test"
}

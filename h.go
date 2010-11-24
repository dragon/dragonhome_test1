package main

type sMyTest struct{
       A int
    }

type i_test interface {
    getValue() int
}

func (s *sMyTest) getValue() int{
        return(s.A)
    }

func (s *sMyTest) test() int{
        return(s.getValue())
    }

type sMyTestOut struct{
        sMyTest
        b int
    }

func (s *sMyTestOut) getValue() int{
        return(s.b+s.sMyTest.getValue())
    }

func test(i i_test){
    println(i.getValue())
}

func (i i_test) test2(){
    println(i.getValue())
}

func main() {
//    println("test12")
    t := new(sMyTestOut)
    t.A = 45
    t.b = 453
    println("")
    println(t.test())
    println(t.getValue())
    test(t)
}




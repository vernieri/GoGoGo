//go build

package main

// Date-Format 

import ( 
    "fmt"
    s "strings"
)
var p = fmt.Println
var nl = fmt.Print

func main() {
   var dia string= "20" //Day 
   var mes string = "08" // Month
   var ano string = "1996" //Year
   
   
   nl("Original: ")
   p(dia, mes, ano)
   nl("Notacao com Barra: ")
   setSlash(dia, mes, ano)
   nl("Notacao com Tracinho: ")
   setTrace(dia, mes, ano)
   nl("Notacao com Demencia: ")
   setImperialistNotation(mes, dia, ano)
   nl("Notacao com Demencia, desta vez com tracinho: ")
   setImperialistNotationWithTrace(mes, dia, ano)
}

func setSlash(d string, m string, a string) string{
    //ano_mes := d[2:len(d)] 
    //p(ano_mes)
    x := s.Join([]string{d, m, a}, "/")
    p(x)
    
    return d
    
}

func setTrace(d string, m string, a string) string{
    x := s.Join([]string{d, m, a}, "-")
    p(x)
    
    return d
}

func setImperialistNotation(m string, d string, a string) string{
    // hur dur american notation
    shit_notation := s.Join([]string{m, d, a}, "/")
    // stop using imperialist trash tho.
    p(shit_notation)
    
    return m
    
}

func setImperialistNotationWithTrace(m string, d string, a string) string{
    // Support Communist propaganda
    shit_notation := s.Join([]string{m, d, a}, "-")
    p(shit_notation)
    
    return m
    
}

// Stay at school, kids.

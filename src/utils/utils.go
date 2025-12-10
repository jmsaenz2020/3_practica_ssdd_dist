package utils

import (
  "fmt"
)

const BOLD = "\033[1;37m"
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const GREEN = "\033[1;32m"
const BLUE = "\033[1;34m"
const END = "\033[0m"

func ErrorMsg(msg string){
  fmt.Printf("%s%s%s\n", RED, msg, END)
}

func WarningMsg(msg string){
  fmt.Printf("%s%s%s\n", YELLOW, msg, END)
}

func InfoMsg(msg string){
  fmt.Printf("%s\n", msg)
}

func BoldMsg(msg string){
  fmt.Printf("%s%s%s\n", BOLD, msg, END)
}

func LeerFecha(aux *string){
  var dia int
  var mes int
  var anyo int

  for{
    fmt.Println("Día")
    LeerInt(&dia)
    fmt.Println("Mes")
    LeerInt(&mes)
    fmt.Println("Año")
    LeerInt(&anyo)
    
    if (dia > 0 && dia <= 31 && mes > 0 && mes <= 12 && anyo > 0){
      *aux = fmt.Sprintf("%d-%d-%d", dia, mes, anyo)
      return
    } else if (dia == 0 && mes == 0 && anyo == 0){
      return
    }
  }
}

func LeerInt(i *int){
  for{
    fmt.Print("> ")
    fmt.Scanf("%d", i)
    if *i >= 0{
      break
    } else {
      WarningMsg("Valor entero inválido")
    }
  }
}

func LeerStr(str *string){
  for{
    fmt.Print("> ")
    fmt.Scanf("%s", str)
    if len(*str) > 0{
      break
    } else {
      WarningMsg("Cadena de texto inválida")
    }
  }
}

func MenuFunc(menu []string) (int, int){
  var opt int

  menu = append(menu, "Salir")
  fmt.Printf("%s%s%s\n", BOLD, menu[0], END) // Menu title

  for i:= 1; i < len(menu); i++{
    fmt.Printf("%d.- %s\n", i, menu[i])
  }

  LeerInt(&opt)

  if opt > 0 && opt < len(menu) - 1{
    return opt, 0
  } else if opt == len(menu) - 1{
    return opt, 2
  }
  return 0, 1
}


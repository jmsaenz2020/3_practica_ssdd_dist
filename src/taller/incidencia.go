package taller

import (
  "fmt"
  "2_practica_ssdd_dist/utils"
)

const TIEMPO_MECANICA = 5
const TIEMPO_ELECTRONICA = 7
const TIEMPO_CARROCERÍA = 11
const MAX_TIEMPO = 15

type Incidencia struct{
  Id int
  Mecanicos []Mecanico
  Tipo int // 1 (Mecánica), 2 (Electrónica), 3(Carrocería)
  Prioridad int // 1 a 3 (Alta a baja)
  Descripcion string
  Estado int // 0 (Cerrado), 1 (Abierta), 2 (En proceso)
}

func (i Incidencia) Info() (string)
{
  return fmt.Sprintf("%s (%03d)", i.Descripcion, i.Id)
}

func (i Incidencia) Visualizar()
{
  fmt.Printf("%sId: %s%03d\n", utils.BOLD, utils.END, i.Id)
  fmt.Printf("%sTipo: %s%d\n", utils.BOLD, utils.END, i.Tipo)
  fmt.Printf("%sPrioridad: %s%d\n", utils.BOLD, utils.END, i.Prioridad)
  fmt.Printf("%sDescripción: %s%s\n", utils.BOLD, utils.END, i.Descripcion)
  fmt.Printf("%sEstado: %s%d\n", utils.BOLD, utils.END, i.Estado)
  utils.BoldMsg("MECÁNICOS")
  if len(i.Mecanicos) > 0{
    for _, m := range i.Mecanicos{
      fmt.Printf("  ·", m.Info())
    }
    fmt.Println()
  } else {
    utils.BoldMsg("SIN MECÁNICOS")
  }
}

func (i *Incidencia) Menu()
{
menu := []string{
  "Menu de incidencia",
  "Visualizar",
  "Modificar"}

for{
  menu[0] = fmt.Sprintf("Menu de %s", i.Info())

  opt, status := utils.MenuFunc(menu)

  if status == 0{
    switch opt{
      case 1:
        i.Visualizar()
      case 2:
        i.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (i *Incidencia) Inicializar()
{
  var exit bool = false

  utils.BoldMsg("ID")
  utils.LeerInt(&i.Id)
  if i.Id == 0{
    exit = true
  }

  if !exit{
    for{
      menu_esp := []string{
        "Selecciona tipo",
        "Mecánica",
        "Electrónica",
        "Carrocería"}
      opt, status := utils.MenuFunc(menu_esp)
      if status == 0{
        i.Tipo = opt - 1
        break
      } else if status == 2{
        exit = true
        break
      }
    }
  }

  if !exit{
  utils.BoldMsg("Prioridad")
    utils.LeerInt(&i.Prioridad)
    if i.Prioridad == 0{
      exit = true
    }
  }

  if !exit{
  utils.BoldMsg("Descripción")
    utils.LeerStr(&i.Descripcion)
    if len(i.Descripcion) == 0{
      exit = true
    } else {
      i.Estado = 1
    }
  }

}

func (i *Incidencia) Modificar()
{
  
  menu := []string{
    "Modificar datos de incidencia",
    "Tipo",
    "Prioridad",
    "Descripción"}
  var buf string
  var num int

  for{
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerInt(&num)
          i.Tipo = num
          utils.InfoMsg("Tipo de incidencia modificado")
        case 2:
          utils.LeerInt(&num)
          i.Prioridad = num
          utils.InfoMsg("Prioridad modificada")
        case 3:
          utils.LeerStr(&buf)
          i.Descripcion = buf
          utils.InfoMsg("Descripción modificada")
      }
    } else if status == 2{
      break
    }
  }
}

func (i Incidencia) Valido() (bool)
{
  return i.Id > 0
}

func (i1 Incidencia) Igual(i2 Incidencia) (bool)
{
  return i1.Id == i2.Id
}

func (i Incidencia) TieneMecanico(m_in Mecanico) (bool)
{
  var tiene bool = false

  for _, m := range i.Mecanicos{
    fmt.Println(m.Info())
    if m.Igual(m_in){
      tiene = true
    }
  }

  return tiene
}

func (i *Incidencia) AsignarMecanico(m Mecanico)
{
  if !i.TieneMecanico(m){
    i.Mecanicos = append(i.Mecanicos, m)
    if i.Estado == 1{
      i.Estado = 2
    }
  } else {
    utils.ErrorMsg("El mecánico ya esta incidencia asignada")
  }
}

package taller

import (
  "fmt"
  "2_practica_ssdd_dist/utils"
)

type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecanica, Electrica, Carroceria
  Experiencia int
  Alta bool
}

func (m *Mecanico)Menu(){
  menu := []string{
    "Menu de mecánico",
    "Visualizar",
    "Modificar"}
  
  for{
    menu[0] = fmt.Sprintf("Menu de %s", m.Nombre)

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          m.Visualizar()
        case 2:
          m.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (m *Mecanico)Inicializar(){
  var exit bool = false  

  utils.BoldMsg("Nombre")
  utils.LeerStr(&m.Nombre)
  if len(m.Nombre) == 0{
    exit = true
  }

  if !exit{
    menu_esp := []string{
    "Selecciona especialidad",
    "Mecánica",
    "Electrónica",
    "Carrocería"}
    for{
      opt, status := utils.MenuFunc(menu_esp)
      if status != 1{
        if status == 0{
          m.Especialidad = opt - 1
        } else {
          exit = true
        }
        break
      }
    }
  }

  if !exit{
    utils.BoldMsg("Experiencia")
    utils.LeerInt(&m.Experiencia)
    m.Id = 1
  }
}

func (m Mecanico)Info() (string){
  return fmt.Sprintf("%s (%03d)", m.Nombre, m.Id)
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%sID: %s%03d\n", utils.BOLD, utils.END, m.Id)
  fmt.Printf("%sNombre: %s%s\n", utils.BOLD, utils.END, m.Nombre)
  fmt.Printf("%sEspecialidad: %s%s\n", utils.BOLD, utils.END, m.ObtenerEspecialidad())
  fmt.Printf("%sExperiencia: %s%d años\n", utils.BOLD, utils.END, m.Experiencia)
  fmt.Printf("%s¿Está de alta? %s%t\n", utils.BOLD, utils.END, m.Alta)
}

func (m Mecanico)Valido() (bool){

  return m.Id > 0 && m.Id <= 999 && len(m.Nombre) > 0 && m.Experiencia >= 0 && m.Especialidad >= 0 && m.Especialidad <= 2
}

func (m1 Mecanico)Igual(m2 Mecanico) (bool){
  return m1.Id == m2.Id
}

func (m *Mecanico)Modificar(){
  menu := []string{
    "Modificar datos de mecánico",
    "Nombre",
    "Especialidad",
    "Experiencia",
    "Dar de baja"}
  var buf string
  var num int

  for{
    if !m.Alta{
      menu[len(menu) - 1] = "Dar de alta"
    } else {
      menu[len(menu) - 1] = "Dar de baja"
    }
    menu[0] = fmt.Sprintf("Modificar datos de %s", m.Nombre)
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerStr(&buf)
          m.Nombre = buf
          utils.InfoMsg("Nombre modificado")
        case 2:
          menu_esp := []string{
            "Selecciona especialidad",
            "Mecánica",
            "Electrónica",
            "Carrocería"}
          opt, status = utils.MenuFunc(menu_esp)
          if status == 0{
            esp := m.ObtenerEspecialidad()
            m.Especialidad = opt - 1
            msg := fmt.Sprintf("Especialidad modificada: %s->%s", esp, m.ObtenerEspecialidad())
            utils.InfoMsg(msg)
          }
        case 3:
          utils.LeerInt(&num)
          m.Experiencia = num
          utils.InfoMsg("Experiencia modificada")
        case 4:
          m.Alta = !m.Alta
          utils.InfoMsg("Estado modificado")
      }
    } else if status == 2{
      break
    }
  }
}

func (m Mecanico)ObtenerEspecialidad() (string){
  switch m.Especialidad{
    case 0:
      return "Mecánica"
    case 1:
      return "Electrónica"
    case 2:
      return "Carrocería"
    default:
      return "Sin especialidad"
  }
}

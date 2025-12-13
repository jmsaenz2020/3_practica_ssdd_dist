package taller

import (
  "fmt"
  "3_practica_ssdd_dist/utils"
)

type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecanica, Electrica, Carroceria
  Experiencia int
  Alta bool
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

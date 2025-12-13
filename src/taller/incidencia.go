package taller

import (
  "fmt"
  "time"
  "3_practica_ssdd_dist/utils"
)

const TIEMPO_ALTA = 5
const TIEMPO_MEDIA = 3
const TIEMPO_BAJA = 1

type Incidencia struct{
  Id int
  Mecanicos []Mecanico
  Tipo int // 1 (Mecánica), 2 (Electrónica), 3(Carrocería)
  Prioridad int // 1 a 3 (Alta a baja)
  Descripcion string
  Estado int // 0 (Cerrado), 1 (Abierta), 2 (En proceso)
}

func (i Incidencia) Info() (string){
  return fmt.Sprintf("%s (%03d)", i.Descripcion, i.Id)
}

func (i Incidencia) Visualizar(){
  fmt.Printf("%sId: %s%03d\n", utils.BOLD, utils.END, i.Id)
  fmt.Printf("%sTipo: %s%d\n", utils.BOLD, utils.END, i.Tipo)
  fmt.Printf("%sPrioridad: %s%d\n", utils.BOLD, utils.END, i.Prioridad)
  fmt.Printf("%sDescripción: %s%s\n", utils.BOLD, utils.END, i.Descripcion)
  fmt.Printf("%sEstado: %s%d\n", utils.BOLD, utils.END, i.Estado)
  utils.BoldMsg("MECÁNICOS")
  if len(i.Mecanicos) > 0{
    for _, m := range i.Mecanicos{
      fmt.Printf("  ·%s", m.Info())
    }
    fmt.Println()
  } else {
    utils.BoldMsg("SIN MECÁNICOS")
  }
}

func (i Incidencia) ObtenerDuracion() (time.Duration){
  var tiempo time.Duration

  switch(i.Tipo){
    case 1:
      tiempo = TIEMPO_ALTA*time.Second
    case 2:
      tiempo = TIEMPO_MEDIA*time.Second
    case 3:
      tiempo = TIEMPO_BAJA*time.Second
  }

  return tiempo
}

func (i Incidencia) Valido() (bool){
  return i.Id > 0 && i.Tipo >= 1 && i.Tipo <= 3 && i.Prioridad >= 1 && i.Prioridad <= 3 && i.Estado >= 0 && i.Estado <= 2 && len(i.Descripcion) > 0
}

func (i1 Incidencia) Igual(i2 Incidencia) (bool){
  return i1.Id == i2.Id
}

func (i Incidencia) TieneMecanico(m_in Mecanico) (bool){
  var tiene bool = false

  for _, m := range i.Mecanicos{
    fmt.Println(m.Info())
    if m.Igual(m_in){
      tiene = true
    }
  }

  return tiene
}

func (i *Incidencia) AsignarMecanico(m Mecanico){
  if !i.TieneMecanico(m){
    i.Mecanicos = append(i.Mecanicos, m)
    if i.Estado == 1{
      i.Estado = 2
    }
  } else {
    utils.ErrorMsg("El mecánico ya esta incidencia asignada")
  }
}

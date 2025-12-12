package taller

import (
  "fmt"
  "time"
  "3_practica_ssdd_dist/utils"
)

const TIEMPO_ESPERA = 15 * time.Second
const NUM_FASES = 4

type Vehiculo struct{
  Matricula int
  Marca string
  Modelo string
  FechaEntrada string
  FechaSalida string
  Incidencia Incidencia
}

func (v Vehiculo) Info() (string){
  return fmt.Sprintf("%s %s (%05d)", v.Marca, v.Modelo, v.Matricula)
}

func (v Vehiculo) Visualizar(){
  fmt.Printf("%sMatricula: %s%05d\n", utils.BOLD, utils.END, v.Matricula)
  fmt.Printf("%sMarca: %s%s\n", utils.BOLD, utils.END, v.Marca)
  fmt.Printf("%sModelo: %s%s\n", utils.BOLD, utils.END, v.Modelo)
  fmt.Printf("%sFecha de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaSalida)
  utils.BoldMsg("Incidencia: ")
  if v.Incidencia.Valido(){
    fmt.Printf("  ·", v.Incidencia.Info())
  } else {
    utils.BoldMsg("SIN INCIDENCIA")
  }
}

func (v *Vehiculo) Menu(){
  menu := []string{
    "Menu de vehiculo",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", v.Info())

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          v.Visualizar()
        case 2:
          //v.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (v *Vehiculo) Inicializar(){
  var exit bool = false
  var inc Incidencia

  utils.BoldMsg("Matrícula")
  utils.LeerInt(&v.Matricula)
  if v.Matricula == 0{
    exit = true
  }

  if !exit{
    utils.BoldMsg("Marca")
    utils.LeerStr(&v.Marca)
    if len(v.Marca) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Modelo")
    utils.LeerStr(&v.Modelo)
    if len(v.Modelo) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Fecha de entrada")
    utils.LeerStr(&v.FechaEntrada)
    if len(v.FechaEntrada) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Fecha de estimada de salida")
    utils.LeerStr(&v.FechaSalida)
    if len(v.FechaSalida) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Incidencia")
    inc.Inicializar()
    v.CrearIncidencia(inc.Tipo, inc.Descripcion)
  }
}
/*
func (v *Vehiculo) Modificar(){

  menu := []string{
    "Modificar datos de vehículo",
    "Matricula",
    "Marca y modelo",
    "Fecha de entrada",
    "Fecha estimada de salida",
    "Incidencias"}
  var buf string
  var num int

  for{
    menu[0] = fmt.Sprintf("Modificar datos de %s", v.Info())
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerInt(&num)
          v.Matricula = num
          utils.InfoMsg("Matricula modificada")
        case 2:
          utils.LeerStr(&buf)
          v.Marca = buf
          utils.LeerStr(&buf)
          v.Modelo = buf
          utils.InfoMsg("Marca y modelo modificado")
        case 3:
          utils.LeerFecha(&v.FechaEntrada)
          utils.InfoMsg("Fecha de entrada modificada")
        case 4:
          utils.LeerFecha(&v.FechaSalida)
          utils.InfoMsg("Fecha estimada de salida modificada")
        case 5:
          v.MenuIncidencias()
      }
    } else if status == 2{
      break
    }
  }
}*/
/*
func (v *Vehiculo) MenuIncidencias(){
  var i Incidencia
  menu := []string{
    "Seleccione una incidencia",
    "Crear incidencia",
    "Eliminar incidencia"}

  for{
    menu = []string{
      "Seleccione una incidencia",
      "Crear incidencia",
      "Eliminar incidencia"}
    for _, i := range v.Incidencias{
      menu = append(menu, i.Info())
    }

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      if opt == 1{
        i.Inicializar()
        v.CrearIncidencia(i.Tipo, i.Descripcion)
      } else if opt == 2{
        i = v.SeleccionarIncidencia()
        v.EliminarIncidencia(i)
      } else {
        v.Incidencias[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}
*/

func (v *Vehiculo) EliminarIncidencia(){
  var i Incidencia
  v.Incidencia = i
}

func (v *Vehiculo)Rutina(t *Taller){
  defer t.Grupo.Done()

  t.AsignarPlaza(v)

  // Fase 1 a 4
  for i := 1; i <= NUM_FASES; i++{
    msg := fmt.Sprintf("Fase %d", i)
    utils.InfoMsg(msg)
    time.Sleep(v.Incidencia.ObtenerDuracion())
  }

  t.SalirVehiculo(v)
}

func (v *Vehiculo) CrearIncidencia(tipo int, descripcion string){
  var i Incidencia

  i.Tipo = tipo
  i.Prioridad = tipo
  i.Descripcion = descripcion
  i.Id = 1

  if i.Valido(){
    v.Incidencia = i
  } else {
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (v Vehiculo) Valido() (bool){
  return v.Matricula > 0 && len(v.Marca) > 0 && len(v.Modelo) > 0 && v.Incidencia.Valido()
}

func (v1 Vehiculo) Igual(v2 Vehiculo) (bool){
  return v1.Matricula == v2.Matricula
}

func (v Vehiculo) StringEstado() (string){
  var estado string

  switch(v.Incidencia.Estado){
    case 0:
      estado = utils.RED
    case 1:
      estado = utils.GREEN
    case 2:
      estado = utils.YELLOW
  }

  return fmt.Sprintf("%s•%s", estado, utils.END)
}


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
  utils.BoldMsg("Incidencias: ")
  if len(v.Incidencias) > 0{
    for _, inc := range v.Incidencias{
      fmt.Printf("  ·", inc.Info())
    }
  } else {
    utils.BoldMsg("SIN INCIDENCIAS")
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
          v.Modificar()
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
}

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

func (v Vehiculo) SeleccionarIncidencia() (Incidencia){
  var i Incidencia

  if len(v.Incidencias) > 0{
    i = v.Incidencias[0]
  }

  return i
}

func (v *Vehiculo) EliminarIncidencia(i Incidencia){

  indice := v.ObtenerIndiceIncidencia(i)
    
  if indice >= 0{ // Eliminar
    lista := v.Incidencias
    lista[indice] = lista[len(lista)-1]
    lista = lista[:len(lista)-1]
    v.Incidencias = lista
  } else {
    utils.ErrorMsg("No se pudo eliminar la incidencia")
  }
}

func (v Vehiculo) ObtenerIncidencia() (Incidencia){
  var inc Incidencia
  
  menu := []string{"Seleccione una incidencia"}

  for{
    menu = []string{"Seleccione una incidencia"}
    for _, i := range v.Incidencias{
      menu = append(menu, i.Info())
    }

    opt, status := utils.MenuFunc(menu)

    if status != 1{
      if status == 0{
        inc = v.Incidencias[opt - 1]
      }
      break
    }
  }

  return inc
}

func (v *Vehiculo)Rutina(t *Taller){
  defer t.Grupo.Done()
  //var fase int = 1

  // Fase 1
  t.EntrarVehiculo(v)
  t.Grupo.Add(1)

  // Fase 2 a 4
  for i := 1; i <= NUM_FASES; i++{
    time.Sleep(v.IncidenciaObtenerDuracion)
  }

  t.SalirVehiculo(v)
}

func (v Vehiculo) ObtenerIndiceIncidencia(i_in Incidencia) (int){
  var res int = -1

  for i, inc := range v.Incidencias{
    if inc.Igual(i_in){
      res = i
    }
  }

  return res
}

func (v *Vehiculo) CrearIncidencia(tipo int, descripcion string){
  var i Incidencia

  i.Tipo = tipo
  i.Prioridad = tipo
  i.Descripcion = descripcion
  i.Id = 1

  if i.Valido() && v.ObtenerIndiceIncidencia(i) == -1{
    v.Incidencias = append(v.Incidencias, i)
  } else {
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (v Vehiculo) Valido() (bool){
  return v.Matricula > 0 && len(v.Marca) > 0 && len(v.Modelo) > 0
}

func (v1 Vehiculo) Igual(v2 Vehiculo) (bool){
  return v1.Matricula == v2.Matricula
}

func (v Vehiculo) StringEstado() (string){
  var estado string = fmt.Sprintf("%s•%s", utils.GREEN, utils.END)
  var cerrado bool = true

  for _, inc := range v.Incidencias{
    if inc.Estado == 2{
      estado = fmt.Sprintf("%s•%s", utils.YELLOW, utils.END)
      return estado
    } else if inc.Estado == 1{
      cerrado = false
    }
  }

  if cerrado{
    estado = fmt.Sprintf("%s•%s", utils.RED, utils.END)
  }

  return estado
}


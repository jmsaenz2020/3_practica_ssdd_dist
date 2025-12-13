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
  return fmt.Sprintf("%s %s (%s)", v.Marca, v.Modelo, v.StringMatricula())
}

func (v Vehiculo) Visualizar(){
  fmt.Printf("%sMatricula: %s%s\n", utils.BOLD, utils.END, v.StringMatricula())
  fmt.Printf("%sMarca: %s%s\n", utils.BOLD, utils.END, v.Marca)
  fmt.Printf("%sModelo: %s%s\n", utils.BOLD, utils.END, v.Modelo)
  fmt.Printf("%sFecha de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaSalida)
  utils.BoldMsg("Incidencia: ")
  if v.Incidencia.Valido(){
    fmt.Printf("  ·%s", v.Incidencia.Info())
  } else {
    utils.BoldMsg("SIN INCIDENCIA")
  }
}

func (v Vehiculo) StringMatricula() (string){
  return fmt.Sprintf("%05d", v.Matricula)
}

func (v *Vehiculo) EliminarIncidencia(){
  var i Incidencia
  v.Incidencia = i // Incidencia vacia
}

func (v Vehiculo)Log(fase int, inicio time.Time){
    tiempo := time.Now()
    msg := fmt.Sprintf("Tiempo %s Coche %s Incidencia %d Fase %d Estado %d", tiempo.Sub(inicio), v.StringMatricula(), v.Incidencia.Id, fase, v.Incidencia.Estado)
    utils.InfoMsg(msg)
}

func (v *Vehiculo)Rutina(t *Taller){
  defer t.Grupo.Done()

  t.AsignarPlaza(v)

  // Fase 1 a 4
  for i := 1; i <= NUM_FASES; i++{
    v.Incidencia.Estado = 2
    v.Log(i, t.TiempoInicio)
    time.Sleep(v.Incidencia.ObtenerDuracion())
    if i == NUM_FASES{
      v.Incidencia.Estado = 0
    } else {
      v.Incidencia.Estado = 1
    }
    v.Log(i, t.TiempoInicio)
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


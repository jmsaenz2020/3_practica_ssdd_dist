package taller

import (
  "fmt"
  "sync"
  "time"
  "3_practica_ssdd_dist/utils"
)

const PLAZAS_MECANICO = 2

type Taller struct{
  Clientes []Cliente
  Plazas chan *Vehiculo
  Mecanicos []Mecanico
  UltimoIdMecanico int
  UltimoIdIncidencia int
  Grupo sync.WaitGroup
  Cerradura sync.RWMutex
  TiempoInicio time.Time
}

func (t *Taller)Inicializar(){
  t.Plazas = make(chan *Vehiculo)
  t.TiempoInicio = time.Now()
}

func (t *Taller)Liberar(){
  t.Cerradura.Lock()
  select{
    case _, ok := <- t.Plazas:
      if !ok && len(t.Plazas) == 0{
        close(t.Plazas)
      }
    default:
      close(t.Plazas)
  }
  t.Cerradura.Unlock()
}

func (t Taller) HayEspacio() (bool){
  vehiculos := t.ObtenerPlazas()

  return len(vehiculos) < PLAZAS_MECANICO*len(t.Mecanicos)
}

func (t *Taller) AsignarPlaza(v *Vehiculo){
  if t.HayEspacio() && v.Valido(){
    t.EntrarVehiculo(v)
  } else if v.Incidencia.Valido(){
    utils.WarningMsg("El vehiculo no tiene una incidencia definida")
  }
}

func (t Taller) Estado(){
  var v Vehiculo
  var i int = 1
  plazas := t.ObtenerPlazas()

  for i, v = range plazas{
    fmt.Printf("%d.- ", i)
    if v.Valido(){
      fmt.Printf("%s %s", v.StringEstado(), v.Info())
    }
    fmt.Println()
  }
}

func (t *Taller) AsignarVehiculo(){
  matriculas := t.ObtenerMatriculaVehiculos()
  var num int
  //var v Vehiculo
  var hayEspacio bool = t.HayEspacio()

  if len(matriculas) > 0 && hayEspacio{
    utils.BoldMsg("VEHICULOS DISPONIBLES")
    for _, m := range matriculas{
      fmt.Println(m)
    }
    fmt.Println("Escriba la matrícula del vehículo a asignar")
    utils.LeerInt(&num)
    //for _, c := range t.Clientes{
      //v = c.ObtenerVehiculoPorMatricula(num)
      //t.AsignarPlaza(v)
    //}
  } else if !hayEspacio{
    utils.WarningMsg("El taller está lleno")
  } else {
    utils.WarningMsg("No hay incidencias en el taller")
  }
}

func (t *Taller) EntrarVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  utils.InfoMsg("entra vehiculo")
  t.Plazas <- v
  t.Cerradura.Unlock()
}

func (t *Taller) SalirVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  <- t.Plazas
  t.Cerradura.Unlock()
}

func (t *Taller)ModificarTaller(){
  taller := make(chan *Vehiculo, PLAZAS_MECANICO*len(t.Mecanicos))

  select{
    case p := <- t.Plazas:
      taller <- p
    default:
      close(taller)
  }

  t.Cerradura.Lock()
  t.Plazas = taller
  t.Cerradura.Unlock()
}

func (t *Taller) CrearMecanico(nombre string, especialidad int, experiencia int){
  var m Mecanico

  m.Nombre = nombre
  m.Especialidad = especialidad
  m.Experiencia = experiencia
  m.Id = t.UltimoIdMecanico + 1

  if m.Valido() && t.ObtenerIndiceMecanico(m) == -1{
    t.UltimoIdMecanico++
    m.Id = t.UltimoIdMecanico
    m.Alta = true
    t.Mecanicos = append(t.Mecanicos, m)
    go t.ModificarTaller()
  } else {
    utils.ErrorMsg("No se ha podido crear al mecanico")
  }
}

func (t *Taller) CrearCliente(c Cliente){
  if c.Valido(){
    t.Clientes = append(t.Clientes, c)
  }
}

func (t *Taller) EliminarMecanico(m Mecanico){
  
  indice := t.ObtenerIndiceMecanico(m)
    
  if indice >= 0{ // Eliminar
    lista := t.Mecanicos
    lista[indice] = lista[len(lista)-1]
    lista = lista[:len(lista)-1]
    t.Mecanicos = lista
  } else {
    utils.ErrorMsg("No se pudo eliminar al mecánico")
  }
}

func (t *Taller) EliminarCliente(c Cliente){
  indice := t.ObtenerIndiceCliente(c)
    
  if indice >= 0{ // Eliminar
    lista := t.Clientes
    lista = lista[:indice+copy(lista[indice:], lista[indice+1:])]
    t.Clientes = lista
  } else {
    utils.ErrorMsg("No se pudo eliminar al mecánico")
  }
}

func (t Taller) ObtenerIndiceMecanico(m_in Mecanico) (int){
  var res int = -1

  for i, m := range t.Mecanicos{
    if m.Igual(m_in){
      res = i
    }
  }

  return res
}

func (t Taller) ObtenerMecanicoPorId(id int) (Mecanico){
  var res Mecanico

  for i, m := range t.Mecanicos{
    if m.Id == id{
      res = t.Mecanicos[i]
    }
  }

  return res
}

func (t Taller) ObtenerClientePorId(id int) (Cliente){
  var res Cliente

  for i, m := range t.Clientes{
    if m.Id == id{
      res = t.Clientes[i]
    }
  }

  return res
}

func (t Taller) ObtenerIndiceCliente(c_in Cliente) (int){
  var res int = -1

  for i, c := range t.Clientes{
    if c.Igual(c_in){
      res = i
    }
  }

  return res
}

func (t Taller) ObtenerMecanicosDisponibles() ([]Mecanico){
  var mecanicos []Mecanico  

  for _, m := range t.Mecanicos{
    if m.Alta{
      mecanicos = append(mecanicos, m)
    }
  }

  return mecanicos
}

func (t Taller) ObtenerIncidencias() ([]Incidencia){
  var incidencias []Incidencia

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      incidencias = append(incidencias, v.Incidencia)
    }
  }

  return incidencias
}

func (t Taller) ObtenerClientesEnTaller() ([]Cliente){
  var clientes []Cliente

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      t.Cerradura.RLock()
      for p := range t.Plazas{
        if v.Igual(*p) && !c.ExisteCliente(clientes){
          clientes = append(clientes, c)
          break // Se ha encontrado el vehiculo del cliente
        }
      }
      t.Cerradura.RUnlock()
    }
  }

  return clientes
}

func (t Taller) ObtenerMatriculaVehiculos() ([]int){
  var matriculas []int

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      if v.Incidencia.Valido(){
        matriculas = append(matriculas, v.Matricula)
      }
    }
  }

  return matriculas
}

func (t Taller) ObtenerPlazas() ([]Vehiculo){
  var vehiculos []Vehiculo
  var v Vehiculo
  var exit bool = false

  for{
    t.Cerradura.RLock()
    select{
      case p := <- t.Plazas:
        v = *p
        if v.Valido(){
          vehiculos = append(vehiculos, v)
        }
        t.Plazas <- p
      default:
        t.Cerradura.RUnlock()
        exit = true
    }
    if exit{
      break
    }
  }

  return vehiculos
}

func (t Taller) ObtenerIncidenciasMecanico(m_in Mecanico) ([]Incidencia){
  var incidencias []Incidencia

  for _, inc := range incidencias{
    if inc.TieneMecanico(m_in){
      incidencias = append(incidencias, inc)
    }
  }

  return incidencias
}

func (t Taller) IncidenciasMecanico(){
  menu := []string{"Seleccione el mecánico"}

  for _, m := range t.Mecanicos{
    menu = append(menu, m.Info())
  }

  for{
    opt, status := utils.MenuFunc(menu)

    if status != 1{
      if status == 0{
        incidencias := t.ObtenerIncidenciasMecanico(t.Mecanicos[opt - 1])
        if len(incidencias) == 0{
          utils.BoldMsg("SIN INCIDENCIAS")
        } else {
          for _, inc := range incidencias{
            fmt.Println(inc.Info())
          }
        }
      }
      break
    }
  }
}

func (t Taller) MecanicosDisponibles(){
  for _, m := range t.Mecanicos{
    if m.Alta{
      fmt.Println(m.Info())
    }
  }
}

func (t *Taller) ModificarMecanico(modif Mecanico){
  for i, m := range t.Mecanicos{
    if m.Igual(modif){
      t.Mecanicos[i] = modif
    }
  }
}

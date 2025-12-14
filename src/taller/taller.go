package taller

import (
  "fmt"
  "sync"
  "time"
  "3_practica_ssdd_dist/utils"
)

const NUM_MECANICOS = 1
const NUM_PLAZAS = 1

type Taller struct{
  Clientes []Cliente
  Plazas chan *Vehiculo
  NumPlazas int
  NumMecanicos int
  Mecanicos []Mecanico
  UltimoIdMecanico int
  UltimoIdIncidencia int
  Grupo sync.WaitGroup
  Cerradura sync.RWMutex
  TiempoInicio time.Time
}

func (t *Taller)Inicializar(){
  t.Plazas = make(chan *Vehiculo, NUM_PLAZAS)
  t.TiempoInicio = time.Now()
}

func (t *Taller)Liberar(){
  t.Grupo.Wait()
  select{
    case _, ok := <- t.Plazas:
      if !ok && len(t.Plazas) == 0{
        close(t.Plazas)
      }
    default:
      close(t.Plazas)
  }
}

func (t Taller) CocheEnTaller(v Vehiculo) (bool){
  var c Cliente
  c.Vehiculos = t.ObtenerPlazas()

  return c.ObtenerIndiceVehiculo(v) >= 0
}

func (t Taller) HayEspacio() (bool){
  vehiculos := t.ObtenerPlazas()

  return len(vehiculos) <= NUM_PLAZAS
}

func (t *Taller) AsignarPlaza(v *Vehiculo) (bool){
  if t.HayEspacio() && v.Valido() && v.Incidencia.Valido(){
    t.EntrarVehiculo(v)
    return true
  } else if !v.Incidencia.Valido(){
    utils.WarningMsg("El vehiculo no tiene una incidencia definida")
  }
  return false
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

func (t *Taller) EntrarVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  t.Plazas <- v
  t.Cerradura.Unlock()
}

func (t *Taller) SalirVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  <- t.Plazas
  t.Cerradura.Unlock()
}

func (t *Taller) AsignarMecanicoAutomatico(v *Vehiculo){
  if t.HayEspacio(){
    for _, m := range(t.ObtenerMecanicosDisponibles()){
      if m.Especialidad == v.Incidencia.Tipo{
        v.Incidencia.AsignarMecanico(m)
      }
    }
  }
}

func (t *Taller) CrearMecanico(nombre string, especialidad int, experiencia int){
  var m Mecanico

  m.Nombre = nombre
  m.Especialidad = especialidad
  m.Experiencia = experiencia
  m.Id = t.UltimoIdMecanico + 1

  if m.Valido(){
    t.UltimoIdMecanico++
    t.Mecanicos = append(t.Mecanicos, m)
  } else{
    utils.ErrorMsg("Mecanico no se ha podido crear")
  }
}

func (t *Taller) CrearCliente(c Cliente){
  if c.Valido(){
    t.Clientes = append(t.Clientes, c)
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

  t.Cerradura.Lock()
  for{
    select{
      case p := <- t.Plazas:
        if p != nil{
          v = *p
          if v.Valido(){
            vehiculos = append(vehiculos, v)
          }
        }
        t.Plazas <- p
      default:
        exit = true
    }
    if exit{
      t.Cerradura.Unlock()
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

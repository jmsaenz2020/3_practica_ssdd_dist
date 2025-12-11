package taller

import (
  "fmt"
  "sync"
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
}

func (t *Taller)Inicializar(){
  t.Plazas = make(chan *Vehiculo)
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

func (t *Taller) MenuPrincipal(){
  menu := []string{
    "Menu principal",
    "Taller",
    "Clientes",
    "Mecánicos"}


  for{
    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          t.Menu()
        case 2:
          t.MenuClientes()
        case 3:
          t.MenuMecanicos()
      }
    } else if status == 2{
      break
    }
  }
}

func (t *Taller) Menu(){
  menu := []string{
    "Menu del taller",
    "Asignar vehiculo",
    "Asignar mecánico",
    "Estado del taller",
    "Listar incidencias",
    "Listar clientes con vehiculos en el taller",
    "Listar incidencias de mecánico",
    "Mecánicos disponibles"}

  for{
    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          t.AsignarVehiculo()
        case 2:
          t.AsignarMecanico()
        case 3:
          t.Estado()
        case 4:
          incidencias := t.ObtenerIncidencias()

          for _, inc := range incidencias{
            fmt.Println(inc.Info())
          }
        case 5:
          clientes := t.ObtenerClientesEnTaller()

          for _, c := range clientes{
            fmt.Println(c.Info())
          }
        case 6:
          if len(t.Mecanicos) > 0{
            t.IncidenciasMecanico()
          } else {
            utils.WarningMsg("No hay mecánicos en el taller")
          }
        case 7:
          t.MecanicosDisponibles()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (t *Taller) MenuMecanicos(){
  var menu []string
  var m Mecanico
  var id int

  for{
    menu = []string{
      "Selecciona un mecánico",
      "Crear Mecánico",
      "Eliminar Mecánico"}
    for _, m := range t.Mecanicos{
      menu = append(menu, m.Info())
    }

    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          m.Inicializar()
          t.CrearMecanico(m.Nombre, m.Especialidad, m.Experiencia)
          if !m.Valido() {
            utils.ErrorMsg("No se ha creado el mecánico")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del mecánico")
            utils.LeerInt(&id)
            m = t.ObtenerMecanicoPorId(id)
            if m.Valido(){
              t.EliminarMecanico(m)
              break
            }
          }
        default:
          t.Mecanicos[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (t Taller) SeleccionarMecanico() (Mecanico){
  var menu []string
  var m Mecanico

  for{
    menu = []string{"Selecciona un mecánico"}
    for _, m := range t.Mecanicos{
      menu = append(menu, m.Info())
    }

    opt, status := utils.MenuFunc(menu)
    
    if status != 1{
      if status == 0{
        m = t.Mecanicos[opt - 1]
      }
      break
    }
  }

  return m
}

func (t *Taller) MenuClientes(){
  var menu []string
  var c Cliente
  var id int

  for{
    menu = []string{
      "Selecciona un cliente",
      "Crear Cliente",
      "Eliminar Cliente"}
    for _, c := range t.Clientes{
      menu = append(menu, c.Info())
    }

    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          c.Inicializar(t)
          t.CrearCliente(c)
          if !c.Valido() {
            utils.ErrorMsg("No se ha creado el cliente")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del cliente")
            utils.LeerInt(&id)
            c = t.ObtenerClientePorId(id)
            if c.Valido(){
              t.EliminarCliente(c)
              break
            }
          }
        default:
          t.Clientes[opt - 3].Menu(t)
      }
    } else if status == 2{
      break
    }
  }
}

func (t Taller) HayEspacio() (bool){
  vehiculos := t.ObtenerPlazas()

  return len(vehiculos) < PLAZAS_MECANICO*len(t.Mecanicos)
}

func (t *Taller) AsignarPlaza(v *Vehiculo){
  if t.HayEspacio() && v.Valido() && len(v.Incidencias) > 0{
    t.Cerradura.RLock()
    select{
      case p := <- t.Plazas:
         if !p.Valido(){
          t.Cerradura.RUnlock()
          t.EntrarVehiculo(v)
          t.Cerradura.RLock()
        }
        t.Plazas <- p
      default:
        t.Cerradura.RUnlock()
        t.EntrarVehiculo(v)
        t.Cerradura.RLock()
    }
    t.Cerradura.RUnlock()
  } else if len(v.Incidencias) == 0{
    utils.WarningMsg("El vehiculo no tiene incidencias que atender")
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
  t.Plazas <- v
  t.Cerradura.Unlock()
  msg := fmt.Sprintf("Vehiculo %s ha entrado al taller", v.Info())
  utils.InfoMsg(msg)
}

func (t *Taller) SalirVehiculo(v *Vehiculo){
  msg := fmt.Sprintf("Vehiculo %s ha salido del taller", v.Info())
  utils.InfoMsg(msg)
  t.Cerradura.Lock()
  <- t.Plazas
  t.Cerradura.Unlock()
}

func (t *Taller) AsignarMecanico(){
  menu := []string{"Seleccione un vehículo"}
  var plaza Vehiculo

  plazas := t.ObtenerPlazas()

  for _, p := range plazas{
    menu = append(menu, p.Info())
  }    

  for {
    if len(menu) > 1{
      opt, status := utils.MenuFunc(menu)

      if status == 0{
        plaza = plazas[opt - 1]
        if len(plaza.Incidencias) > 0{
          inc := plaza.ObtenerIncidencia()
          if inc.Valido(){
            m := t.SeleccionarMecanico()
            if m.Valido(){
              inc.AsignarMecanico(m)
            }
          }
        } else {
          utils.WarningMsg("El vehiculo no tiene incidencias")
        }
      } else if status == 2{
        break
      }
    } else {
      utils.WarningMsg("No hay vehículos en el taller")
      break
    }
  }
}

func (t *Taller)ModificarTaller(){
  taller := make(chan *Vehiculo, PLAZAS_MECANICO*len(t.Mecanicos))

  select{
    case p := <- t.Plazas:
      taller <- p
    default:
      close(taller)
  }

  t.Plazas = taller
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
      for _, inc := range v.Incidencias{
        incidencias = append(incidencias, inc)
      }
    }
  }

  return incidencias
}

func (t Taller) ObtenerClientesEnTaller() ([]Cliente){
  var clientes []Cliente

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      t.Cerradura.Lock()
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
      if len(v.Incidencias) > 0{
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
        exit = true
    }
    if exit{
      t.Cerradura.RUnlock()
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

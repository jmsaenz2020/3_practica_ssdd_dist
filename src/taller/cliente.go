package taller

import (
  "fmt"
  "3_practica_ssdd_dist/utils"
)

type Cliente struct{
  Id int
  Nombre string
  Telefono int
  Email string
  Vehiculos []Vehiculo
}

const MAX_ID_CLIENTE = 100000000 - 1
const MAX_TELEFONO_CLIENTE = 1000000000 - 1


func (c Cliente) Info() (string){
  return fmt.Sprintf("%s (%08d)", c.Nombre, c.Id)
}

func (c Cliente) Visualizar(){
  fmt.Printf("%sID: %s%08d\n", utils.BOLD, utils.END, c.Id)
  fmt.Printf("%sNombre: %s%s\n", utils.BOLD, utils.END, c.Nombre)
  fmt.Printf("%sTeléfono: %s%09d\n", utils.BOLD, utils.END, c.Telefono)
  fmt.Printf("%sEmail: %s%s\n", utils.BOLD, utils.END, c.Email)
  fmt.Printf("%sVehiculos:%s\n", utils.BOLD, utils.END)
  c.ListarVehiculos()
}

func (c *Cliente) CrearVehiculo(v Vehiculo, t *Taller){
  if v.Valido() && c.ObtenerIndiceVehiculo(v) == -1{
    c.Vehiculos = append(c.Vehiculos, v)
    if v.Incidencia.Valido(){
      t.Grupo.Add(1)
      go c.Vehiculos[len(c.Vehiculos) - 1].Rutina(t)
    }
  } else {
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (c *Cliente) ModificarIncidenciaVehiculo(v Vehiculo, tipo int, desc string, t *Taller){
  var indice int = c.ObtenerIndiceVehiculo(v)

  if v.Valido() && indice >= 0{
    c.Vehiculos[indice].ModificarIncidencia(tipo, desc)
    t.Grupo.Add(1)
    go c.Vehiculos[indice].Rutina(t)
  } else {
    utils.ErrorMsg("El vehiculo no existe o no está en propiedad del cliente")
  }
}

func (c Cliente) SeleccionarVehiculo() (Vehiculo){
  var v Vehiculo  

  if len(c.Vehiculos) > 0{
    v = c.Vehiculos[0]
  }

  return v
}

func (c *Cliente) EliminarVehiculo(v Vehiculo){

  indice := c.ObtenerIndiceVehiculo(v)
    
  if indice >= 0{ // Eliminar
    lista := c.Vehiculos
    lista[indice] = lista[len(lista)-1]
    lista = lista[:len(lista)-1]
    c.Vehiculos = lista
  } else {
    utils.ErrorMsg("No se pudo eliminar al vehículo")
  }
}

func (c Cliente) ListarVehiculos(){
  if len(c.Vehiculos) > 0{
    for _, v := range c.Vehiculos{
      fmt.Printf("  %s·%s%s\n", utils.BOLD, utils.END, v.Info())
    }
  } else {
    utils.BoldMsg("SIN VEHICULOS")
  }
}

func (c Cliente) Valido() (bool){
  return c.Id > 0 && c.Id <= MAX_ID_CLIENTE && len(c.Nombre) > 0 && c.Telefono > 0  && c.Telefono <= MAX_TELEFONO_CLIENTE && len(c.Email) > 0
}

func (c1 Cliente) Igual(c2 Cliente) (bool){
  return c1.Id == c2.Id
}

func (c Cliente) ObtenerIndiceVehiculo(v_in Vehiculo) (int){
  var res int = -1

  for i, v := range c.Vehiculos{
    if v.Igual(v_in){
      res = i
    }
  }

  return res
}

func (c Cliente) ObtenerVehiculoPorMatricula(matricula int) (Vehiculo){
  var res Vehiculo  

  for _, v := range c.Vehiculos{
    if v.Matricula == matricula{
      res = v
    }
  }

  return res
}

func (c_in Cliente) ExisteCliente(clientes []Cliente) (bool){
  var existe bool = false

  for _, c := range clientes{
    if c.Igual(c_in){
      existe = true
    }
  }

  return existe
}


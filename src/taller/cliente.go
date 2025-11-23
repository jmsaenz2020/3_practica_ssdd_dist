package taller

import (
  "fmt"
  "time"
  "2_practica_ssdd_dist/utils"
)

type Cliente struct{
  Id int
  Nombre string
  Telefono int
  Email string
  Vehiculos []Vehiculo
}

func (c *Cliente) Inicializar()
{
  var exit bool = false

  utils.BoldMsg("ID")
  utils.LeerInt(&c.Id)
  if c.Id == 0{
    exit = true
  }

  if !exit{
    utils.BoldMsg("Nombre")
    utils.LeerStr(&c.Nombre)
    if len(c.Nombre) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Telefono")
    utils.LeerInt(&c.Telefono)
    if c.Telefono == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Email")
    utils.LeerStr(&c.Email)
    if len(c.Email) == 0{
      exit = true
    }
  }

  // Ya está creado el cliente base (sin vehículos)
  if !exit{
    c.MenuVehiculos()
  }
}

func (c Cliente) Info() (string)
{
  return fmt.Sprintf("%s (%08d)", c.Nombre, c.Id)
}

func (c Cliente) Visualizar()
{
  fmt.Printf("%sID: %s%08d\n", utils.BOLD, utils.END, c.Id)
  fmt.Printf("%sNombre: %s%s\n", utils.BOLD, utils.END, c.Nombre)
  fmt.Printf("%sTeléfono: %s%09d\n", utils.BOLD, utils.END, c.Telefono)
  fmt.Printf("%sEmail: %s%s\n", utils.BOLD, utils.END, c.Email)
  fmt.Printf("%sVehiculos:%s\n", utils.BOLD, utils.END)
  c.ListarVehiculos()
}

func (c *Cliente) MenuVehiculos()
{
  var v Vehiculo  
  menu := []string{
    "Seleccione un vehículo",
    "Crear vehículo",
    "Eliminar vehículo"}

  for{
    menu = []string{
      "Seleccione un vehículo",
      "Crear vehículo",
      "Eliminar vehículo"}
    for _, v := range c.Vehiculos{
      menu = append(menu, v.Info())
    }

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      if opt == 1{
        v.Inicializar()
        c.CrearVehiculo(v)
      } else if opt == 2{
        v = c.SeleccionarVehiculo()
        if v.Valido(){
          c.EliminarVehiculo(v)
        }
      } else {
        c.Vehiculos[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (c *Cliente) CrearVehiculo(v Vehiculo)
{
  v.TiempoAcumulado = make(chan time.Duration, 1)
  v.TiempoAcumulado <- 0

  if v.Valido() && c.ObtenerIndiceVehiculo(v) == -1{
    c.Vehiculos = append(c.Vehiculos, v)
  } else {
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (c Cliente) SeleccionarVehiculo() (Vehiculo)
{
  var v Vehiculo  

  if len(c.Vehiculos) > 0{
    v = c.Vehiculos[0]
  }

  return v
}

func (c *Cliente) EliminarVehiculo(v Vehiculo)
{

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

func (c Cliente) ListarVehiculos()
{
  if len(c.Vehiculos) > 0{
    for _, v := range c.Vehiculos{
      fmt.Printf("  %s·%s%s\n", utils.BOLD, utils.END, v.Info())
    }
  } else {
    utils.BoldMsg("SIN VEHICULOS")
  }
}

func (c *Cliente) Menu()
{
  menu := []string{
    "Menu de cliente",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", c.Nombre)

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          c.Visualizar()
        case 2:
          c.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (c *Cliente) Modificar()
{
  menu := []string{
    "Modificar datos de cliente",
    "ID",
    "Nombre",
    "Teléfono",
    "Email",
    "Vehiculos"}
  var buf string
  var num int

  for{
    menu[0] = fmt.Sprintf("Modificar datos de %s", c.Nombre)
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerInt(&num)
          c.Id = num
          utils.InfoMsg("ID modificado")
        case 2:
          utils.LeerStr(&buf)
          c.Nombre = buf
          utils.InfoMsg("Nombre modificado")
        case 3:
          utils.LeerInt(&num)
          c.Telefono = num
          utils.InfoMsg("Teléfono modificado")
        case 4:
          utils.LeerStr(&buf)
          c.Email = buf
          utils.InfoMsg("Email modificado")
        case 5:
          c.MenuVehiculos()
      }
    } else if status == 2{
      break
    }
  }
}

func (c Cliente) Valido() (bool)
{
  return c.Id > 0 && len(c.Nombre) > 0 && c.Telefono > 0 && len(c.Email) > 0
}

func (c1 Cliente) Igual(c2 Cliente) (bool)
{
  return c1.Id == c2.Id
}

func (c Cliente) ObtenerIndiceVehiculo(v_in Vehiculo) (int)
{
  var res int = -1

  for i, v := range c.Vehiculos{
    if v.Igual(v_in){
      res = i
    }
  }

  return res
}

func (c Cliente) ObtenerVehiculoPorMatricula(matricula int) (Vehiculo)
{
  var res Vehiculo  

  for _, v := range c.Vehiculos{
    if v.Matricula == matricula{
      res = v
    }
  }

  return res
}

func (c_in Cliente) ExisteCliente(clientes []Cliente) (bool)
{
  var existe bool = false

  for _, c := range clientes{
    if c.Igual(c_in){
      existe = true
    }
  }

  return existe
}




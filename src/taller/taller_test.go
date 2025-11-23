package taller

import (
  "testing"
)

func TestCopia(test *testing.T){
  var taller Taller
  taller.CrearMecanico("Pepe", 0, 0)
  c := Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  i := Incidencia{Id: 1, Tipo: 1, Prioridad: 1, Descripcion: "Luna delantera rota", Estado: 1}
  v.CrearIncidencia(i)
  c.CrearVehiculo(v)
  v = Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  c.CrearVehiculo(v)
  taller.CrearCliente(c)
  for _, v := range c.Vehiculos{
    taller.AsignarPlaza(v)
  }
  
  if len(taller.ObtenerMatriculaVehiculos()) != 1{
    test.Errorf("ERROR: No se pudo detectar la copia")
  }
}

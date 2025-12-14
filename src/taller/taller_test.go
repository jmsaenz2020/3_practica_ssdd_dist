package taller

import (
  "testing"
)

func ValoresBasicos() (Taller){
  var taller Taller
  taller.Inicializar()
  taller.CrearMecanico("Pepe", 0, 0)
  c := Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  v.CrearIncidencia(1, "Luna delantera rota")
  c.CrearVehiculo(v, &taller)
  v = Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  c.CrearVehiculo(v, &taller)
  taller.CrearCliente(c)

  return taller
}

func TestBasicos(test *testing.T){
  taller := ValoresBasicos()
  taller.Liberar()
}

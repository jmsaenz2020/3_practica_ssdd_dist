package main

import (
  "2_practica_ssdd_dist/taller"
)

func main(){
  var t taller.Taller
  
  // INICIALIZAR
  t.Inicializar()
  t.CrearMecanico("Pepe", 0, 0)
  c := taller.Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := taller.Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  i := taller.Incidencia{Id: 1, Tipo: 1, Prioridad: 1, Descripcion: "Luna delantera rota", Estado: 1}
  i.AsignarMecanico(t.Mecanicos[0])
  v.CrearIncidencia(i)
  c.CrearVehiculo(v)
  v = taller.Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  c.CrearVehiculo(v)
  t.CrearCliente(c)
  t.AsignarPlaza(c.Vehiculos[0])
  t.AsignarPlaza(c.Vehiculos[1])
  // FIN INICIALIZAR

  t.MenuPrincipal()
}

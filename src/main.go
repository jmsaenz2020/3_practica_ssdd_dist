package main

import (
  "3_practica_ssdd_dist/taller"
)

func main(){
  var t taller.Taller

  // INICIALIZAR
  t.Inicializar()
  t.CrearMecanico("Pepe", 0, 0)
  t.CrearMecanico("Pepe", 0, 0)
  c := taller.Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := taller.Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  v.CrearIncidencia(1, "Luna delantera rota")
  v.Incidencia.AsignarMecanico(t.Mecanicos[0])
  c.CrearVehiculo(v, &t)
  v = taller.Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  v.CrearIncidencia(1, "Luna delantera rota")
  v.Incidencia.AsignarMecanico(t.Mecanicos[0])
  c.CrearVehiculo(v, &t)
  t.CrearCliente(c)
  // FIN INICIALIZAR

    
  t.Grupo.Wait()
  t.Liberar()
}

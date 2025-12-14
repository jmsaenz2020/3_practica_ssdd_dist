package main

import (
  "3_practica_ssdd_dist/taller"
  "fmt"
  "math/rand"
)

const sims = 1
const NUM_COCHES = 1
const NUM_MECANICOS = 1

func generarDatosCliente() (taller.Cliente){
  var c taller.Cliente

  c.Id = rand.Int()%taller.MAX_ID_CLIENTE + 1
  c.Telefono = rand.Int()%taller.MAX_TELEFONO_CLIENTE + 1
  c.Nombre = "cliente"
  c.Email = fmt.Sprintf("%s@mail.com", c.Nombre)

  return c
}

func valoresBasicos() (taller.Taller){
  var t taller.Taller
  var v taller.Vehiculo
  var c taller.Cliente

  t.Inicializar()
  for n := 0; n < NUM_MECANICOS; n++{
    t.CrearMecanico("Mecanico", 0, 0)
  }
  c = generarDatosCliente()

  for n := 0; n < NUM_COCHES; n++{
    v = taller.Vehiculo{Matricula: rand.Int()%taller.MAX_MATRICULA + 1, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
    v.ModificarIncidencia(rand.Int()%3 + 1, "Incidencia")
    v.Incidencia.AsignarMecanico(t.Mecanicos[0])
    c.CrearVehiculo(v, &t)
  }
  t.CrearCliente(c)

  return t
}

func main(){
  var t taller.Taller
  
  for i := 0; i < sims; i++{
    t = valoresBasicos()
    t.Liberar()
  }
}
package taller

import(
	"sync"
)

type Plaza struct{
	Vehiculo *Vehiculo
	Estado sync.RWMutex
}
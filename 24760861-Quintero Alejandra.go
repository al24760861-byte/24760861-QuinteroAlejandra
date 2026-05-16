package main

import (
	"fmt"
)

// Definimos la estructura Materia
type Materia struct {
	Nombre       string
	Calificacion int
}

// Lista de materias
var materias []Materia

// Función para saber si está aprobado (>= 70)
func estaAprobado(calificacion int) bool {
	return calificacion >= 70
}

// Función para buscar una materia por nombre
func buscarMateria(nombre string) *Materia {
	for i := range materias {
		if materias[i].Nombre == nombre {
			return &materias[i]
		}
	}
	return nil
}

// Función para eliminar materia
func eliminarMateria(nombre string) {
	for i := range materias {
		if materias[i].Nombre == nombre {
			materias = append(materias[:i], materias[i+1:]...)
			fmt.Println("Materia eliminada:", nombre)
			return
		}
	}
	fmt.Println("Materia no encontrada:", nombre)
}

// Función para actualizar calificación
func actualizarCalificacion(nombre string, nuevaCalificacion int) {
	m := buscarMateria(nombre)
	if m != nil {
		m.Calificacion = nuevaCalificacion
		fmt.Println("Calificación actualizada de", nombre, "a", nuevaCalificacion)
	} else {
		fmt.Println("Materia no encontrada:", nombre)
	}
}

// Función para mostrar materia con mayor calificación
func materiaMayorCalificacion() *Materia {
	if len(materias) == 0 {
		return nil
	}
	mayor := materias[0]
	for _, m := range materias {
		if m.Calificacion > mayor.Calificacion {
			mayor = m
		}
	}
	return &mayor
}

func main() {
	// Agregamos materias de ejemplo
	materias = append(materias, Materia{"Matemáticas", 85})
	materias = append(materias, Materia{"Historia", 60})
	materias = append(materias, Materia{"Programación", 95})

	// Probamos funciones
	historia := buscarMateria("Historia")
	if historia != nil {
		fmt.Println("¿Historia aprobada?", estaAprobado(historia.Calificacion))
	}

	actualizarCalificacion("Historia", 75)

	eliminarMateria("Matemáticas")

	mayor := materiaMayorCalificacion()
	if mayor != nil {
		fmt.Println("Materia con mayor calificación:", mayor.Nombre, "-", mayor.Calificacion)
	}
}

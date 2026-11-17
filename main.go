//Estudiante: David Torres

package main

import (
	"fmt"
	"strings"
)

const estudiante = "David Torres"

type Notificador interface {
	Enviar (destino string) error
	Canal () string
}

type ErrorEnvio struct {
	Canal string
	Destino string
	Motivo string
}

func (e ErrorEnvio) Error() string {
	return e.Canal + ", " + e.Destino + ", " + e.Motivo
}

type Correo struct {}

func (c Correo) Enviar(destino string) error {
	if !strings.Contains(destino, "@"){
		return ErrorEnvio {Canal: "Correo", Destino: destino, Motivo: "falta del @"}
	}
	return nil
}

func (c Correo) Canal () string {
	return "Correo"
}

type SMS struct {}

func (s SMS) Enviar(destino string) error{
	if len(destino) != 10 {
		return ErrorEnvio {Canal: "SMS", Destino: destino, Motivo: "Debe tener 10 digitos"}
	}
	return nil
}

func (s SMS) Canal () string {return "SMS"}

type Cita struct {
	Hora string
	Paciente string
	Contacto string
	Medio Notificador
}

func (c Cita) String() string{
	return c.Hora + "" + c.Paciente + " (" + c.Medio.Canal()+")" 
}

var citas = []Cita{
	{Hora: "8:30", Paciente: "Ana Paredes", Contacto: "ana.paredes@correo.ec", Medio:Correo{}},
	{Hora: "9:00", Paciente: "Luis Guamán", Contacto: "0985669258", Medio:SMS{} },
	{Hora: "9:30", Paciente: "Carlos Ortega", Contacto: "carlos.ortega@correo.ec", Medio:Correo{} },
	{Hora: "10:00", Paciente: "Sofía Chicaiza", Contacto: "0982997725", Medio:SMS{} },
}

func costo(n Notificador) float64{
	switch n.(type) {
		case Correo:
			return 0.00
		case SMS:
			return 0.05
		default:
			return 0
	}
}

func main() {
	fmt.Println("Recordatorios del Consultorio Carolina")
	fmt.Println("Realizado por:", estudiante)
	fmt.Println()

	enviados := 0
	costoTotal := 0.0

	for _, c := range citas {
		err := c.Medio.Enviar(c.Contacto)
		if err != nil {
			fmt.Printf("%s - error: %v\n", c, err)
			continue
		}
		fmt.Printf("%s - enviado\n", c)
		enviados++
		costoTotal += costo(c.Medio)
	}

	fmt.Println()
	fmt.Printf("Enviados: %d de %d\n", enviados, len(citas))
	fmt.Printf("Costo de los envíos: $%.2f\n", costoTotal)
}
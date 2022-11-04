package helpers

import "time"

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Docente struct {
	ID               int       `json:"ID"`
	Nombre           string    `json:"Nombre"`
	Apellido         string    `json:"Apellido"`
	Matricula        string    `json:"Matricula"`
	Fecha_Nacimiento time.Time `json:"fecha_nacimiento"`
	Titulo           string    `json:"titulo"`
	Correo           string    `json:"correo"`
	Telefono         string    `json:"telefono"`
}

type GetDocentesResponseStruct struct {
	Docentes []*Docente `json:"docentes"`
}

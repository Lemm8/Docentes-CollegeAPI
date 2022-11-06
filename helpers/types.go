package helpers

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Docente struct {
	ID               int    `json:"ID"`
	Nombre           string `json:"Nombre"`
	Apellido         string `json:"Apellido"`
	Matricula        string `json:"Matricula"`
	Fecha_Nacimiento string `json:"fecha_nacimiento"`
	Titulo           string `json:"titulo"`
	Correo           string `json:"correo"`
	Telefono         string `json:"telefono"`
}

type ListDocentesResponseStruct struct {
	Status   int        `json:"status"`
	Docentes []*Docente `json:"docentes"`
}

type DocenteResponseStruct struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Docente *Docente `json:"docente"`
}

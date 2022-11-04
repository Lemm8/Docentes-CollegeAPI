package helpers

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Docente struct {
	Nombre           string `json:"Nombre"`
	Apellido         string `json:"Apellido"`
	Matricula        string `json:"Matricula"`
	Fecha_Nacimiento string `json:"fecha_nacimiento"`
	Titulo           string `json:"titulo"`
	Correo           string `json:"correo"`
	Telefono         string `json:"telefono"`
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ServicioPotencia interface {
	Calcular(base, exponente int) int
}

type potenciaImpl struct{}

func (p potenciaImpl) Calcular(base, exponente int) int {
	if exponente < 0 {
		return 0
	}
	resultado := 1
	for i := 0; i < exponente; i++ {
		resultado *= base
	}
	return resultado
}

type Respuesta struct {
	Base      int `json:"base"`
	Exponente int `json:"exponente"`
	Resultado int `json:"resultado"`
}

type Controlador struct {
	servicio ServicioPotencia
}

func (c Controlador) Manejar(w http.ResponseWriter, r *http.Request) {
	base, _ := strconv.Atoi(r.URL.Query().Get("base"))
	exp, _ := strconv.Atoi(r.URL.Query().Get("exp"))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Respuesta{base, exp, c.servicio.Calcular(base, exp)})
}

func main() {
	ctrl := Controlador{servicio: potenciaImpl{}}
	http.HandleFunc("/potencia", ctrl.Manejar)
	log.Println("API en :8080 - GET /potencia?base=2&exp=3")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

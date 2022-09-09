package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Modulocpu = "/proc/cpu_201709311"

var conn = MySQLConnection()

func MySQLConnection() *sql.DB {
	usuario := "root"
	pass := "Guatemala2022"
	host := "tcp(34.122.120.202:3306)"
	db := "tarea2"
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, db))
	if err != nil {
		fmt.Println("HAY ERROR: \n", err)
	} else {
		fmt.Println("se ha conectado a mysql!")
	}
	return conn
}

//struct para cpu
type Cpu struct {
	Procesos []struct {
		Pid         int    `json:"Pid"`
		Nombre      string `json:"Nombre"`
		Estado      int    `json:"Estado"`
		User        int    `json:"User"`
		Mem         int    `json:"Mem"`
		Subprocesos []struct {
			Pid    int    `json:"Pid"`
			Nombre string `json:"Nombre"`
			Ppid   int    `json:"Ppid"`
		} `json:"Subprocesos"`
	} `json:"Procesos"`
	Estados []struct {
		Ejecucion  int `json:"Ejecucion"`
		Suspendido int `json:"Suspendido"`
		Detenido   int `json:"Detenido"`
		Zombie     int `json:"Zombie"`
	} `json:"Estados"`
}

var C Cpu

func getModuloCpu(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(Modulocpu)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &C)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(C)
}

func createProceso(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	borrar := "DELETE FROM Procesos"
	_, err := conn.Exec(borrar)
	if err != nil {
		fmt.Println("Error al borrar")
	}
	b, err := json.Marshal(C)
	if err != nil {
		fmt.Println("Error al convertir a json")
	}
	query := "INSERT INTO Procesos VALUES('" + string(b) + "');"
	result, err := conn.Exec(query)
	if err != nil {
		fmt.Println("Error al insertar")
	}

	json.NewEncoder(response).Encode(result)
}

func getProcesos(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var procesos []Cpu
	query := "SELECT * FROM Procesos"
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Println("Error al consultar")
	}
	for rows.Next() {
		var proceso Cpu
		var b []byte
		err = rows.Scan(&b)
		if err != nil {
			fmt.Println("Error al escanear")
		}
		err = json.Unmarshal(b, &proceso)
		if err != nil {
			fmt.Println("Error al convertir a json")
		}
		procesos = append(procesos, proceso)
	}
	json.NewEncoder(response).Encode(procesos)
}

func main() {
	fmt.Println("Servidor corriendo en el puerto 8080")
	router := mux.NewRouter()
	router.HandleFunc("/cpu", getModuloCpu).Methods("GET")
	router.HandleFunc("/cpuinsert", createProceso).Methods("POST")
	router.HandleFunc("/cpuget", getProcesos).Methods("GET")
	http.ListenAndServe(":8080", router)
}

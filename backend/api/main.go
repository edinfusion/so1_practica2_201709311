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
	for i := 0; i < len(C.Procesos); i++ {
		query := "INSERT INTO Subprocesos(Pid, Nombre, Ppid) VALUES (?,?,?)"
		for j := 0; j < len(C.Procesos[i].Subprocesos); j++ {
			_, err := conn.Exec(query, C.Procesos[i].Subprocesos[j].Pid, C.Procesos[i].Subprocesos[j].Nombre, C.Procesos[i].Subprocesos[j].Ppid)
			if err != nil {
				fmt.Println("HAY ERROR: \n", err)
			}
		}
		query = "INSERT INTO Procesos(Pid, Nombre, Estado, User, Mem) VALUES (?,?,?,?,?)"
		_, err := conn.Exec(query, C.Procesos[i].Pid, C.Procesos[i].Nombre, C.Procesos[i].Estado, C.Procesos[i].User, C.Procesos[i].Mem)
		if err != nil {
			fmt.Println("HAY ERROR: \n", err)
		}

	}
	json.NewEncoder(response).Encode(C)
}

func main() {
	fmt.Println("Servidor corriendo en el puerto 8080")
	router := mux.NewRouter()
	router.HandleFunc("/cpu", getModuloCpu).Methods("GET")
	router.HandleFunc("/cpuinsert", createProceso).Methods("POST")
	http.ListenAndServe(":8080", router)
}

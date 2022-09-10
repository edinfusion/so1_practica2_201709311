package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mackerelio/go-osstat/cpu"
)

var Modulocpu = "/proc/cpu_201709311"
var Moduloram = "/proc/ram_201709311"

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

type Ram struct {
	Porcentaje string `json:"Porcentaje"`
}

type Cpuso struct {
	Porcentaje string `json:"Porcentaje"`
}

var C Cpu
var R Ram
var B Cpuso

func getModuloCpu() {
	data, err := ioutil.ReadFile(Modulocpu)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &C)
	if err != nil {
		fmt.Println(err)
	}
}

func getModuloRam() {
	data, err := ioutil.ReadFile(Moduloram)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &R)
	if err != nil {
		fmt.Println(err)
	}
}

func cpuUsage() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(100) * time.Millisecond)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	//fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
	uso := 100 - (float64(after.Idle-before.Idle) / total * 100)
	a := strconv.FormatFloat(uso, 'f', 2, 64)
	fmt.Println(a)
	B.Porcentaje = a
}

func createProceso(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	getModuloCpu()
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

func createRam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	getModuloRam()
	borrar := "DELETE FROM Ram"
	_, err := conn.Exec(borrar)
	if err != nil {
		fmt.Println("Error al borrar")
	}
	b, err := json.Marshal(R)
	if err != nil {
		fmt.Println("Error al convertir a json")
	}
	query := "INSERT INTO Ram VALUES('" + string(b) + "');"
	result, err := conn.Exec(query)
	if err != nil {
		fmt.Println("Error al insertar")
	}

	json.NewEncoder(response).Encode(result)
}

func createUsoCpu(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	cpuUsage()
	borrar := "DELETE FROM Cpu"
	_, err := conn.Exec(borrar)
	if err != nil {
		fmt.Println("Error al borrar")
	}
	b, err := json.Marshal(B)
	if err != nil {
		fmt.Println("Error al convertir a json")
	}
	query := "INSERT INTO Cpu VALUES('" + string(b) + "');"
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

func getRam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var rams []Ram
	query := "SELECT * FROM Ram"
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Println("Error al consultar")
	}
	for rows.Next() {
		var ram Ram
		var b []byte
		err = rows.Scan(&b)
		if err != nil {
			fmt.Println("Error al escanear")
		}
		err = json.Unmarshal(b, &ram)
		if err != nil {
			fmt.Println("Error al convertir a json")
		}
		rams = append(rams, ram)
	}
	json.NewEncoder(response).Encode(rams)
}

func main() {
	cpuUsage()
	fmt.Println("Servidor corriendo en el puerto 8080")
	router := mux.NewRouter()
	router.HandleFunc("/cpuinsert", createProceso).Methods("POST")
	router.HandleFunc("/raminsert", createRam).Methods("POST")
	router.HandleFunc("/usocpuinsert", createUsoCpu).Methods("POST")
	router.HandleFunc("/cpuget", getProcesos).Methods("GET")
	http.ListenAndServe(":8080", router)
}

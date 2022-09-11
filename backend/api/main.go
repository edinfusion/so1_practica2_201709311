package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

type EstadosI struct {
	Estados []struct {
		Ejecucion  int `json:"Ejecucion"`
		Suspendido int `json:"Suspendido"`
		Detenido   int `json:"Detenido"`
		Zombie     int `json:"Zombie"`
	} `json:"Estados"`
}

type ProcesosI struct {
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
}

type Ram struct {
	Porcentaje string `json:"Porcentaje"`
}

type Cpuso struct {
	Porcentaje string `json:"Porcentaje"`
}

//var C Cpu
var R Ram
var B Cpuso
var E EstadosI
var P ProcesosI

func getModuloCpu() {
	data, err := ioutil.ReadFile(Modulocpu)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &P)
	if err != nil {
		fmt.Println(err)
	}
}

func getEstados() {
	data, err := ioutil.ReadFile(Modulocpu)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &E)
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

func createProcesos() {
	getModuloCpu()
	borrar := "DELETE FROM Procesos"
	_, err := conn.Exec(borrar)
	if err != nil {
		fmt.Println("Error al borrar")
	}
	b, err := json.Marshal(P)
	if err != nil {
		fmt.Println("Error al convertir a json")
	}
	query := "INSERT INTO Procesos VALUES('" + string(b) + "');"
	_, err1 := conn.Exec(query)
	if err1 != nil {
		fmt.Println("Error al insertar")
	}
	fmt.Println("Procesos agregados")
}

func createEstados() {
	getEstados()
	borrar := "DELETE FROM Estados"
	_, err := conn.Exec(borrar)
	if err != nil {
		fmt.Println("Error al borrar")
	}
	b, err := json.Marshal(E)
	if err != nil {
		fmt.Println("Error al convertir a json")
	}
	query := "INSERT INTO Estados VALUES('" + string(b) + "');"
	_, err1 := conn.Exec(query)
	if err1 != nil {
		fmt.Println("Error al insertar")
	}
	fmt.Println("Estados actualizados")
}

func createRam() {
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
	_, err1 := conn.Exec(query)
	if err1 != nil {
		fmt.Println("Error al insertar")
	}
	fmt.Println("Ram actualizada")
}

func createUsoCpu() {
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
	_, err1 := conn.Exec(query)
	if err1 != nil {
		fmt.Println("Error al insertar")
	}
	fmt.Println("Cpu actualizado")
}

func main() {
	fmt.Println("SERVIDOR GO INSERTANDO MODULOS EN MYSQL GCP")
	escuchar := true
	createProcesos()
	for escuchar {
		fmt.Println("Escuchando Modulos...")
		createEstados()
		time.Sleep(time.Duration(2000) * time.Millisecond)
		createRam()
		time.Sleep(time.Duration(2000) * time.Millisecond)
		createUsoCpu()
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
}

////struct para cpu
//type Cpu struct {
//	Procesos []struct {
//		Pid         int    `json:"Pid"`
//		Nombre      string `json:"Nombre"`
//		Estado      int    `json:"Estado"`
//		User        int    `json:"User"`
//		Mem         int    `json:"Mem"`
//		Subprocesos []struct {
//			Pid    int    `json:"Pid"`
//			Nombre string `json:"Nombre"`
//			Ppid   int    `json:"Ppid"`
//		} `json:"Subprocesos"`
//	} `json:"Procesos"`
//	Estados []struct {
//		Ejecucion  int `json:"Ejecucion"`
//		Suspendido int `json:"Suspendido"`
//		Detenido   int `json:"Detenido"`
//		Zombie     int `json:"Zombie"`
//	} `json:"Estados"`
//}

//func createProcesos(response http.ResponseWriter, request *http.Request) {
//	response.Header().Add("content-type", "application/json")
//	getModuloCpu()
//	borrar := "DELETE FROM Procesos"
//	_, err := conn.Exec(borrar)
//	if err != nil {
//		fmt.Println("Error al borrar")
//	}
//	b, err := json.Marshal(P)
//	if err != nil {
//		fmt.Println("Error al convertir a json")
//	}
//	query := "INSERT INTO Procesos VALUES('" + string(b) + "');"
//	result, err := conn.Exec(query)
//	if err != nil {
//		fmt.Println("Error al insertar")
//	}
//	json.NewEncoder(response).Encode(result)
//	fmt.Println("Se inserto correctamente")
//}

//func createEstados(response http.ResponseWriter, request *http.Request) {
//	response.Header().Add("content-type", "application/json")
//	getEstados()
//	borrar := "DELETE FROM Estados"
//	_, err := conn.Exec(borrar)
//	if err != nil {
//		fmt.Println("Error al borrar")
//	}
//	b, err := json.Marshal(E)
//	if err != nil {
//		fmt.Println("Error al convertir a json")
//	}
//	query := "INSERT INTO Estados VALUES('" + string(b) + "');"
//	result, err := conn.Exec(query)
//	if err != nil {
//		fmt.Println("Error al insertar")
//	}
//	json.NewEncoder(response).Encode(result)
//}

//func createRam(response http.ResponseWriter, request *http.Request) {
//	response.Header().Add("content-type", "application/json")
//	getModuloRam()
//	borrar := "DELETE FROM Ram"
//	_, err := conn.Exec(borrar)
//	if err != nil {
//		fmt.Println("Error al borrar")
//	}
//	b, err := json.Marshal(R)
//	if err != nil {
//		fmt.Println("Error al convertir a json")
//	}
//	query := "INSERT INTO Ram VALUES('" + string(b) + "');"
//	result, err := conn.Exec(query)
//	if err != nil {
//		fmt.Println("Error al insertar")
//	}
//
//	json.NewEncoder(response).Encode(result)
//}

//func createUsoCpu(response http.ResponseWriter, request *http.Request) {
//	response.Header().Add("content-type", "application/json")
//	cpuUsage()
//	borrar := "DELETE FROM Cpu"
//	_, err := conn.Exec(borrar)
//	if err != nil {
//		fmt.Println("Error al borrar")
//	}
//	b, err := json.Marshal(B)
//	if err != nil {
//		fmt.Println("Error al convertir a json")
//	}
//	query := "INSERT INTO Cpu VALUES('" + string(b) + "');"
//	result, err := conn.Exec(query)
//	if err != nil {
//		fmt.Println("Error al insertar")
//	}
//
//	json.NewEncoder(response).Encode(result)
//}

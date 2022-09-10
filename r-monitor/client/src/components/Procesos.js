import React, { useState, useEffect } from 'react';
import { Container, Card, Table } from 'react-bootstrap';
import axios from 'axios';
import Proceso from './Proceso';
import Subproceso from './Subproceso';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Estadop from './Estados';


const Procesos = () => {

    const [procesos, setProcesos] = useState([]);
    const [estados, setEstados] = useState([]);
    const [refreshData, setRefreshData] = useState(false);

    useEffect(() => {
        obtenerProcesos();
    }, [])
    if (refreshData) {
        setRefreshData(false);
        obtenerProcesos();
    }



    return (
        <div>
            <Navbar bg="dark" variant="dark" expand="lg">
                <Container fluid>
                    <Navbar.Brand href="#">MONITOR SO1-P2</Navbar.Brand>
                    <Navbar.Toggle aria-controls="navbarScroll" />
                    <Navbar.Collapse id="navbarScroll">
                        <Nav
                            className="me-auto my-2 my-lg-0"
                            style={{ maxHeight: '100px' }}
                            navbarScroll
                        >
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            <br></br>
            <Container>
                {estados != null && estados.map((iproceso) => (
                    <Estadop estate={iproceso} total={procesos.length}/>
                ))}
            </Container>
            <Container>
                <h2>Listado de Procesos</h2>
                <Table responsive hover>
                    <thead className='table-dark' >
                        <tr>
                            <th>ID</th>
                            <th>NOMBRE</th>
                            <th>ESTADO</th>
                            <th>USUARIO</th>
                            <th>MEMORIA</th>
                            <th>Subprocesos</th>
                        </tr>
                    </thead>
                    <tbody>
                        {procesos != null && procesos.map((proceso, i) => (
                                <Proceso iproceso={proceso} sbproceso={proceso.Subprocesos} />
                        ))}

                    </tbody>
                </Table>
            </Container>
        </div>
    );


    //obtenerProcesos();
    //function getProcesos(){
    //    axios.get('/prueba')
    //    .then(res => {
    //        console.log(res.data);
    //    })
    //}
    function obtenerProcesos() {
        var url = '/prueba';
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if (response.status === 200) {
                setProcesos(response.data.Procesos);
                setEstados(response.data.Estados);
                console.log(response.data.Procesos);
            }
        }
        )
    }

}
export default Procesos;
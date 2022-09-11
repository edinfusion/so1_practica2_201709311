import React from 'react';
import { Container, Card, Table,Row } from 'react-bootstrap';
import axios from 'axios';
import Proceso from './Proceso';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Estadop from './Estados';
import Cpu from './Cpu';
import Ram from './Ram';



export default class Procesos extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            procesos: [],
            ram: [],
            cpu: [],
            cpuu: {},
            estados: [],
        }
        this.obtenerProcesos();
    }
    componentDidMount() {
        this.updateProcesos();
    }

    obtenerProcesos = async () => {
        var url = '/prueba';
        await axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if (response.status === 200) {
                this.setState({ procesos: response.data.Procesos });
                //console.log(response.data.Procesos);
                console.log("SIGO OBTENIENDO PROCESOS");
            }
        }
        )
    }



    //const [procesos, setProcesos] = useState([]);
    //const [ram,setRam] = useState([]);
    //const [usocpu,setUsoCpu] = useState([]);
    //
    //const [estados, setEstados] = useState([]);
    //const [refreshData, setRefreshData] = useState(false);


    updateProcesos = async () => {
        function sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        }

        try {
            while (true) {

                // estados(response.data.Estados);
                //setRefreshData(true);
                await sleep(2000);
                var url = '/cpu';
                const res = await axios.get(url);
                this.setState({ cpuu: res.data });
                console.log(this.state.cpuu);
                await sleep(2000);
                var url = '/estados';
                const res2 = await axios.get(url);
                this.setState({ estados: res2.data.Estados });
                await sleep(2000);
                var url = '/ram';
                const res3 = await axios.get(url);
                this.setState({ ram: res3.data });
            }
        }
        catch (error) {
            console.error("Ocurrio un error, ya no se pudo actualizar");
            console.log(error);
        }

    };


    render() {
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
                <Row>
                    
                        {this.state.estados != null && this.state.estados.map((iproceso) => (
                            <Estadop estate={iproceso} total={this.state.procesos.length} />
                        ))}
                        <Cpu uso={this.state.cpuu} />
                        <Ram usor={this.state.ram} />
                  
                </Row>
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
                            {this.state.procesos != null && this.state.procesos.map((proceso, i) => (
                                <Proceso iproceso={proceso} sbproceso={proceso.Subprocesos} />
                            ))}
                        </tbody>
                    </Table>
                </Container>

            </div>
        );




        //async function obtenerProcesos() {
        //    var url = '/prueba';
        //    await axios.get(url, {
        //        responseType: 'json'
        //    }).then(response => {
        //        if (response.status === 200) {
        //            setProcesos(response.data.Procesos);
        //            setEstados(response.data.Estados);
        //            //console.log(response.data.Procesos);
        //        }
        //    }
        //    )
        //}
        //async function obtenerRam() {
        //    var url = '/ram';
        //    await axios.get(url, {
        //        responseType: 'json'
        //    }).then(response => {
        //        if (response.status === 200) {
        //            setRam(response.data);
        //            //console.log(response.data);
        //        }
        //    }
        //    )
        //}
        //
        //async function obtenerUsoCpu() {
        //    var url = '/cpu';
        //    await axios.get(url, {
        //        responseType: 'json'
        //    }).then(response => {
        //        if (response.status === 200) {
        //            setUsoCpu(response.data);
        //            //console.log(response.data);
        //        }
        //    }
        //    )
        //}


    }

}
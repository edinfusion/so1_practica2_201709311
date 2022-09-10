import 'bootstrap/dist/css/bootstrap.css';
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/js/src/collapse.js";
import Collapsible from 'react-collapsible';
import Subproceso from './Subproceso';
const Proceso = ({ iproceso, sbproceso }) => {
    return (
        <tr>
            <td>{iproceso.Pid != undefined && iproceso.Pid}</td>
            <td>{iproceso.Nombre != undefined && iproceso.Nombre}</td>
            <td>{iproceso.Estado != undefined && iproceso.Estado}</td>
            <td>{iproceso.User != undefined && iproceso.User}</td>
            <td>{iproceso.Mem != undefined && iproceso.Mem}</td>
            <td><Collapsible trigger="Click aqui para ver subprocesos">
            <table className="table table-striped table-dark">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>NOMBRE</th>
                    </tr>
                </thead>
                <tbody>
                    {sbproceso != null && sbproceso.map((subproceso, i) => (
                        <Subproceso isubproceso={subproceso} />
                    ))}
                </tbody>
            </table>
            </Collapsible>
            </td>
        </tr>
        
    );
    //return (
    //    <Collapsible trigger={
    //        <tr>
    //            <td>{iproceso.Pid != undefined && iproceso.Pid}</td>
    //            <td>{iproceso.Nombre != undefined && iproceso.Nombre}</td>
    //            <td>{iproceso.Estado != undefined && iproceso.Estado}</td>
    //            <td>{iproceso.User != undefined && iproceso.User}</td>
    //            <td>{iproceso.Mem != undefined && iproceso.Mem}</td>
    //        </tr>
    //    }>
    //        <table className="table table-striped table-dark">
    //            <thead>
    //                <tr>
    //                    <th>ID</th>
    //                    <th>NOMBRE</th>
    //                </tr>
    //            </thead>
    //            <tbody>
    //                {sbproceso != null && sbproceso.map((subproceso, i) => (
    //                    <Subproceso isubproceso={subproceso} />
    //                ))}
    //            </tbody>
    //        </table>
    //    </Collapsible>
    //);
}

export default Proceso;


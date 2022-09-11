


const Subproceso = ({isubproceso}) => {
    return(
        <tr>
            <td>{isubproceso != undefined && isubproceso.Pid}</td>
            <td>{isubproceso != undefined && isubproceso.Nombre}</td>
        </tr>
    );
}

export default Subproceso;




const Estadop = ({estate,total}) => {
    return(
        <div className="col-md-3 p-4">
            <div className="card">
                <div className="card-body">
                    <h3 className="card-title">Resumen Procesos</h3>
                    <p className="card-text">     Ejecucion: {estate.Ejecucion}</p>
                    <p className="card-text">    Suspendido: {estate.Suspendido}</p>
                    <p className="card-text">      Detenido: {estate.Detenido}</p>
                    <p className="card-text">        Zombie: {estate.Zombie}</p>
                    <p className="card-text">Total Procesos: {total}</p>
                </div>
            </div>
        </div>
    );
}

export default Estadop;



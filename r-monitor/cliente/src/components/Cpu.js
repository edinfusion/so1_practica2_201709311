
import {Chart as ChartJs,Tooltip,Title,ArcElement,Legend} from 'chart.js';
import {Pie} from 'react-chartjs-2';
ChartJs.register(
    Tooltip,Title,ArcElement,Legend
);

const usageCpu = ({uso}) => {
const data = {
    labels: ['Uso','Libre'],
    datasets: [
        {
            label: 'Uso de CPU',
            data: [uso.Porcentaje,100-uso.Porcentaje],
            backgroundColor: [
                'red',
                'green'
            ],
            borderColor: [
                'Black',
                'Gray'
            ],
            borderWidth: 1,
        },
    ],
};
                    

    return(
        <div className="col-md-3 p-4">
            <div className="card">
                <div className="card-body">
                    <h3 className="card-title">Uso de CPU </h3>
                    <p className="card-text">{uso.Porcentaje} %</p>
                </div>
                <Pie data={data} />
            </div>
        </div>
    );
}

export default usageCpu;
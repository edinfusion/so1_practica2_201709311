import {Chart as ChartJs,Tooltip,Title,ArcElement,Legend} from 'chart.js';
import {Pie} from 'react-chartjs-2';
ChartJs.register(
    Tooltip,Title,ArcElement,Legend
);

const usageRam = ({usor}) => {
const data = {
    labels: ['Uso','Libre'],
    datasets: [
        {
            label: 'Uso de RAM',
            data: [usor.Porcentaje,100-usor.Porcentaje],
            backgroundColor: [
                'Blue',
                'LightBlue'
            ],
            borderColor: [
                'Black',
                'Brown'
            ],
            borderWidth: 1,
        },
    ],
};
    return(
        <div className="col-md-3 p-4">
            <div className="card">
                <div className="card-body">
                    <h3 className="card-title">Uso de Ram </h3>
                    <p className="card-text">{usor.Porcentaje} %</p>
                </div>
                <Pie data={data} />
            </div>
        </div>
    );
}

export default usageRam;
var mysql = require('mysql');

const {promisify} = require('util');

const pool = mysql.createPool(
    {
        connectionLimit: 10,
        host: '34.122.120.202',
        user: 'root',
        password: 'Guatemala2022',
        database: 'tarea2',
        port: '3306'
    }
)

pool.getConnection((err, connection) => {
    if (err) {
        if (err.code === 'PROTOCOL_CONNECTION_LOST') {
            console.error('Database connection was closed.')
        }
        if (err.code === 'ER_CON_COUNT_ERROR') {
            console.error('Database has to many connections')
        }
        if (err.code === 'ECONNREFUSED'){
            console.error('DATABASE CONNECTION WAS REFUSED')
        }
    }
    if (connection) connection.release();
console.log('DB is Connected');
return;
});

pool.query = promisify(pool.query);
module.exports = pool;



//var conn = mysql.createConnection({
//    host: '34.122.120.202',
//    user: 'root',
//    password: 'Guatemala2022',
//    database: 'tarea2',
//    port: '3306'
//});
//conn.connect(function(err) {
//    if (err) throw err;
//    console.log("Connected!");
//});
//module.exports = conn;


//func MySQLConnection() *sql.DB {
//	usuario := "root"
//	pass := "Guatemala2022"
//	host := "tcp(34.122.120.202:3306)"
//	db := "tarea2"
//	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, db))
//	if err != nil {
//		fmt.Println("HAY ERROR: \n", err)
//	} else {
//		fmt.Println("se ha conectado a mysql!")
//	}
//	return conn
//}
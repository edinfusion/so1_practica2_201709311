//const { query } = require('express');
var express = require('express'); // Express web server framework
//var conn = require('./db'); 
const pool = require('./db') // DB connection
var cors = require('cors'); // CORS middleware
const port = process.env.PORT || 5000;  // Port number
const app = express();

app.use(cors(
    {
        origin: '*',
    }
));
app.listen(port, () => console.log(`Listening on port ${port}`)); // Start server

app.get('/prueba', async (req, res) => {
    const query = await pool.query('SELECT * FROM Procesos');
    //res.send(query[0].Cpu);
    if (query && query.length > 0){
        res.send(query[0].Cpu);
    }else{
        res.send("...");
        console.log("No hay datos");
    }
});

app.get('/estados', async (req, res) => {
    const query = await pool.query('SELECT * FROM Estados');
    if (query && query.length > 0){
        res.send(query[0].Estados);
    }else{
        res.send("...");
        console.log("No hay datos");
    }
});

app.get('/ram', async (req, res) => {
    const query = await pool.query('SELECT * FROM Ram');
    if (query && query.length > 0){
        res.send(query[0].Ram)
    }else{
        res.send("...");
        console.log("No hay datos");
    }
});

app.get('/cpu', async (req, res) => {
    const query = await pool.query('SELECT * FROM Cpu');
    if (query && query.length > 0){
        res.send(query[0].Cpuu);
    }else{
        res.send("...");
        console.log("No hay datos");
    }
});

//app.get('/prueba', async (req, res) => {
//    var query = "SELECT * FROM Procesos";
//    conn.query(query, function (err, result, fields) {
//        if (err) throw err;
//        res.send(result[0].Cpu);
//    });
//});
//
//app.get('/estados', async (req, res) => {
//    var query = "SELECT * FROM Estados";
//    conn.query(query, function (err, result, fields) {
//        if (err) throw err;
//        res.send(result[0].Estados);
//    });
//});
//
//app.get('/ram', async (req, res) => {
//
//    var query = "SELECT * FROM Ram";
//    conn.query(query, function (err, result, fields) {
//        if (err) throw err;
//        res.send(result[0].Ram);
//    });
//});
//
//app.get('/cpu', async (req, res) => {
//
//    var query = "SELECT * FROM Cpu";
//    conn.query(query, function (err, result, fields) {
//        if (err) throw err;
//        res.send(result[0].Cpuu);
//    });
//});
//
//
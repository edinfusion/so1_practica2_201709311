const { query } = require('express');
var express = require('express'); // Express web server framework
var conn = require('./db'); // DB connection
var cors = require('cors'); // CORS middleware
const port = process.env.PORT || 5000;  // Port number
var app = express(); // Create express app

app.use(cors( 
    {
        origin: '*',
    }
));
app.listen(port, () => console.log(`Listening on port ${port}`)); // Start server


app.get('/prueba',function(req, res){

    conn.query(
        'SELECT * FROM Procesos;',
        function(err, result){
            if (err) throw err;
            //res.json({answer:42});
            res.send(result[0].Cpu);
        }
    );
    
});
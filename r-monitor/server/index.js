const { query } = require('express');
var express = require('express');
var conn = require('./db');
var app = express();
app.use(express.json());


//const Cpu ={
//    Procesos:[
//        {
//            Pid:0,
//            Nombre:"",
//            Estado:0,
//            Usuario:0,
//            Mem:0,
//            Subprocesos:[
//                {
//                    Pid:0,
//                    Nombre:"",
//                    Ppid:0,
//                }
//            ]
//        }
//    ],
//    Estados:[
//        {
//            Ejecucion:0,
//            Suspendido:0,
//            Detenido:0,
//            Zombie:0,
//        }
//    ]
//}

//app.get('/getproces',function(req, res){
//    query = "SELECT * FROM Procesos";
//    rows = conn.query(query, function(err, result){
//        if(err){
//            res.send(err);
//        }else{
//            res.send(result);
//        }
//    });
//    for (let i = 0; i < rows.length; i++) {
//        const element = rows[i];
//        Cpu.Procesos[i].Pid=element.Pid;
//        Cpu.Procesos[i].Nombre=element.Nombre;
//        Cpu.Procesos[i].Estado=element.Estado;
//        Cpu.Procesos[i].Usuario=element.Usuario;
//        Cpu.Procesos[i].Mem=element.Mem;
//    }
//    console.log(Cpu.Procesos);
//});

var cpu = {
    Procesos: [],
    Estados: []
};

app.get('/getproces',function(req, res){

    var rows = conn.query(
        'SELECT * FROM Procesos;',
        function(err, result){
            if (err) throw err;
            //res.json({answer:42});
            res.send(result[0].Cpu);
        }
    );
    
});



app.listen(
    8080, 
    ()=>console.log('Server running on port 8080')
);
    
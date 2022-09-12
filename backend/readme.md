# Pasos para levantar servidor correctamente

1. ingresar a carpeta modules y generar los modulos correspondientes con los comandos:

   make all 

   sudo insmod <nombremodulo>.ko

   esto se hace para el modulo del cpu y ram (en sus respectivas carpetas)

   se verifica si estan insertados con cd /proc y visualizar si ya estan los nombres de cada modulo en la lista 

2. levantar servidor de golang localmente o en contenedor.

   LOCAL:

   ​	go run main.go

   CONTENEDOR:

   ​	docker pull edinmv/backend_go_p2  (para descargar imagen)

   ​	docker images (validar que ya esta descargada, se debe observar el nombre de la imagen)

   ​	docker run -it -p 8080:8080 idcontenedor (para levantar contenedor y correr la app)

3. Si todo salió bien se debe osbservar por terminal:

   ![imagenserveronline](backend/images/servergo.png)
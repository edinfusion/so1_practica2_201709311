# Módulos Kernel en Linux

​		Un módulo del Kernel es un fragmento de código o binarios que pueden ser cargado y eliminados del Kernel según las necesidades de este. Tienen el objetivo de extender sus funcionalidades, son fragmentos de código que pueden ser cargados y eliminados del núcleo.

​		Partiendo de esta base decimos que los módulos del Kernel son fragmentos de código, abierto o cerrado, que nosotros podemos añadir o quitar del kernel con el fin de añadir o quitar una funcionalidad.

​		Al crear módulos o insertar módulos, se pueden registrar temperaturas de componentes de la PC, hacer funcionar determinado dispositivo del Hardware, obtener procesos en ejecución, porcentajes de uso de cada dispositivo conectado al Hardware.

# Modulo CPU

​		Este modulo lista los procesos que se ejecutan en el sistema operativo, en este caso es la distribución de Linux Ubuntu. Para ello se utilizan las siguientes librerías:

- linux/module.h: Para la creación de un modulo siempre es necesario incluir esta libreria, ya que esta librería ofrece bastantes macros que seran necesarias para la programacion de módulos
- linux/kernel.h: Este modulo se incluye para poder ver los mensajes en terminal, a traves de la macro KERN_INFO que para decirlo de una manera general, es el println en terminal.
- linux/init.h: Esta libreria ofrece las macros para inicio y fin,  que es lo primero que lee el kernel, en este caso se tiene como inicio la siguiente estructura, module_init(_insert), estas es la primer macro que se ejecuta con la cual se inserta el modulo, para la macro de salida se tiene module_exit(_remove), la cual elimina el modulo y ejecutar el metodo correspondiente
- linux/hugetlb.h: Esta librería tiene las macros relacionadas a los procesos que se están ejecutando en el CPU, y es asi como se pueden listar todos los procesos.
- linux/profc_fs.h: Esta librería se utiliza para poder acceder al struct procs, ya que al acceder a las propiedades de este struct, se pueden obtener id, memoria, estados de cada proceso
- linux/seq_file.h: Esta libreria ayuda a ir almacenando en el archivo indicado, la informacion que necesitamos para el modulo
- linux/sched/signal.h: Esta libreria nos permite hacer el recorrido para cada proceso (for each proccess).

## Task Struct

a través de la librería indicada (profc_fs) se accede a la estructura llamada task_struct, la cual contiene un proceso en especifico, y se obtienen las diversas propiedades de cada proceso. Las propiedades de este Struct son las siguientes:

- pid: contiene el id del proceso.
- comm: contiene el nombre del proceso
- state: el estado del proceso (este estado es de tipo numero el cual tiene un significado dependiendo del valor) en este caso se utilizan los siguientes valores como referencia:
  - state == 0 : Proceso en ejecución
  - state == 1 : Proceso Suspendido
  - state == 128 : Proceso detenido
  - state == 260: Proceso Zombie

## Métodos para realizar modulo

### static int al_abrir

​		Este metodo crea un struct, en el cual se almacena la información del modulo y redirige hacia la función en la cual se escribe cada proceso.

### static in _insert

​		Este método genera un kern info a la hora de ingresar el modulo a linux, ademas de ser el método disparador hacia las demás funciones.

### static _remove

​		Este método genera un kern info a la hora de que se elimina el modulo de linux.		

Todos estos metodos o funciones se pueden encontrar en la carpeta de módulos:

[VER METODOS MODULO CPU](backend/modules/modulocpu/cpu_201709311.c)

# Modulo RAM

​		Este modulo genera la información sobre la memoria principal del sistema, para poder evaluar así el desempeño y utilización en tiempo real de esta memoria, de acuerdo a los proceso en ejecución. Para ello se utilizan las siguientes librerías (algunas librerias del modulo de cpu se reutilizan en este modulo por lo que no se explican):

- linux/seq_file: Contiene macros para lectura secuencial, seguimiento y escritura secuencial.

- sysinfo: Este struct es por medio del cual se puede obtener la información acerca de la memoria ram.

  - struct sysinfo

    se utilizan los atributos:

    totalram: este atributo se obtiene el tamaño de almacenamiento que posee la memoria ram del sistema

    freeram: El espacio libre que tiene la memoria ram

Este modulo es el mas sencillo y los métodos que se utilizaron fueron similares al del modulo del CPU.

Todos estos metodos o funciones se pueden encontrar en la carpeta de módulos:

[VER METODOS MODULO RAM](backend/modules/moduloram/ram_201709311.c)

# Pasos para insertar modulos:

1. ## Crear archivo MAKE 

Este archivo facilita la escritura de un conjunto de reglas complejas sobre como compilar un programa, reduciendo este trabajo a simplemente escibir MAKE y la regla que se desea ejecutar en el archivo.

En el cuerpo de este archivo se inscriben un nombre para el modulo y dos comandos.

- obj-m: acá se define el nombre del archivo que se desea compilar
- all: este comando tiene los parámetros make - C $(KDIR) M=$(PWD) modules, y genera los archivos correspondientes para crear el modulo, escribiéndolo en un formato de bajo nivel como lo es ASM.
- clean: este comando elimina los archivos compilados que se generan por el comando all.
- Make all: este comando se ejecuta desde consola para crear los archivos correspondientes para crear el modulo
- Make clean: este comando se ejecuta para eliminar los archivos que genera el comando make all

2. ## Insertar modulo

​	Luego de que se generaron los archivos con el comando make all, se procede a insertar el modulo con el comando **insmod** con el archivo .ko que se genera 	con make all. Este comando tiene la siguiente sintaxis:

​							**sudo insmod <nombredearchivo>.ko**

3. ## Descargar modulo

   Cuando ya no es necesario el modulo se puede eliminar con el comando **rmmod** con el archivo .ko que se genera con make all. Este comando tiene la siguiente sintaxis:

   ​					**sudo rmmod <nombredearchivo>.ko**



​                                                                                                  Edin Montenegro - Estudiante de Ingeniería en Ciencias y Sistemas. Curso Sistema Operativos 1.

#include <linux/module.h> /* Todos los modulos lo necesitan */
#include <linux/kernel.h>  /* Ofrece la macro KERN_INFO (los print en terminal) */
#include <linux/init.h>    /* Ofrece las macros de inicio y fin */
#include <linux/hugetlb.h> /* Ofrece la funcion get_cpu() */
#include <linux/proc_fs.h> /* struct procs*/
#include <asm/uaccess.h>	/* copy_from_user() */
#include <linux/seq_file.h> /* seq_printf() */
#include <linux/sched.h> /* struct task_struct */
#include <linux/sched/signal.h> /* for_each_process() */

#define modulo_cpu "cpu_201709311"

/*DOCUMENTACION DEL MODULO*/
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("MODULO CPU");
MODULE_AUTHOR("Edin Montenegro");

struct task_struct *procesos, *subprocesos;
struct list_head *listap;
int cont;
int cont2;
//para estados
long int ejecucion;
long int suspendido;
long int detenido;
long int zombie;

static int listar_procesos(struct seq_file *m, void *v) {
    long procesomem;
    cont2 = 0;
    ejecucion = 0;
    suspendido = 0;
    detenido = 0;
    zombie = 0;

    seq_printf(m,"{\n");
    seq_printf(m, "\"Procesos\":[\n");
    for_each_process(procesos) {
        if(procesos->mm) {
            procesomem = get_mm_rss(procesos->mm);
        }
        if (cont2 == 0){
            cont2 = 1;
        }else{
            seq_printf(m, ",");
        }
        
        //Aqui empiezan los estados
        /*  Procesos en ejecución Número de procesos en ejecución (running)
        Procesos suspendidos Número de procesos suspendidos (sleeping)
        Procesos detenidos Número de procesos detenidos (stopped)
        Procesos zombie Número de procesos zombie
        Total de procesos Número total de procesos*/

        //se ejecuta cada proceso y se revisa su estado
        if(procesos->__state == 0){
            ejecucion++;
        }else if(procesos->__state == 1){
            suspendido++;
        }else if(procesos->__state == 128){
            detenido++;
        }else if(procesos->__state == 260){
            zombie++;
        }
        /*
        PID Identificador del proceso
        Nombre Nombre del proceso
        Usuario Usuario que ejecutó el proceso
        Estado Estado en el que se encuentra el proceso
        %RAM Porcentaje de utilización de RAM por el proceso//se ejecuta cada proceso y se revisa su estado
        */
        seq_printf(m, "\n{ \"PID\" : %d, \"Nombre\" : \"%s\", \"Estado\" : %ld , \"User\" : %i, \"Mem\"  : %ld,", procesos->pid, procesos->comm, procesos->__state, procesos->cred->uid.val, procesomem);
        seq_printf(m, "\"subprocesos\" : [");
        cont = 0;

        list_for_each(listap, &(procesos->children)) {
             if (cont == 0){
                cont = 1;
            }else{
                seq_printf(m, ",");
            }

            subprocesos= list_entry(listap, struct task_struct, sibling);
            seq_printf(m, "\n{ \"PID\" : %d, \"Nombre\" : \"%s\"}", subprocesos->pid, subprocesos->comm);
            if(subprocesos->__state == 0){
                ejecucion++;
            }else if(subprocesos->__state == 1){
                suspendido++;
            }else if(subprocesos->__state == 128){
                detenido++;
            }else if(subprocesos->__state == 260){
                zombie++;
            }

        }

        cont = 0;
        seq_printf(m, "]\n}\n");
    }
    seq_printf(m, "\n],\n");
    //Aqui empiezan los estados

    seq_printf(m,"\"Estados\": [\n{\"Ejecucion\": %li, \"Suspendido\": %li, \"Detenido\": %li, \"Zombie\":%li}\n]\n", ejecucion, suspendido, detenido, zombie);
    seq_printf(m, "}\n");


    return 0;
}
static int al_abrir(struct inode *inode, struct file *file) {
    return single_open(file,listar_procesos,NULL);
}

static struct proc_ops operaciones= {
    .proc_open= al_abrir,
    .proc_read= seq_read 
};

static int _insert(void) {
    printk(KERN_INFO "Edin Montenegro\n");
    proc_create(modulo_cpu,0,NULL, &operaciones);
    return 0;
}

static void _remove(void) {
    printk(KERN_INFO "Segundo Semestre 2022\n");
    remove_proc_entry(modulo_cpu,NULL);
}

module_init(_insert);
module_exit(_remove);
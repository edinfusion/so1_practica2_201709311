
#include <linux/proc_fs.h>
#include <linux/seq_file.h> 
#include <linux/hugetlb.h>
#include <asm/uaccess.h> 
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>   
#include <linux/fs.h>




#define modulo_ram "ram_201709311"

/*DOCUMENTACION DEL MODULO*/
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("MODULO RAM");
MODULE_AUTHOR("Edin Montenegro");
struct sysinfo sysi;
static int uso_ram(struct seq_file *m, void *v) {
    
    si_meminfo(&sysi);
 

    seq_printf(m, "{\"Porcentaje\":\"%d",(((sysi.totalram)-(sysi.freeram))*100)/(sysi.totalram));
    //seq_printf(archivo, "%d",porcentaje);
    seq_printf(m, "\"}");

    return 0;
}


static int al_abrir(struct inode *inode, struct file *file) {
    return single_open(file,uso_ram,NULL);
}

static struct proc_ops operaciones= {
    .proc_open= al_abrir,
    .proc_read= seq_read 
};

static int _insert(void) {
    printk(KERN_INFO "201709311\n");
    proc_create(modulo_ram,0,NULL, &operaciones);
    return 0;
}

static void _remove(void) {
    printk(KERN_INFO "Sistemas Operativos 1\n");
    remove_proc_entry(modulo_ram,NULL);
}

module_init(_insert);
module_exit(_remove);
cmd_/home/edinmv/go/src/Practica_2/backend/modules/moduloram/Module.symvers := sed 's/\.ko$$/\.o/' /home/edinmv/go/src/Practica_2/backend/modules/moduloram/modules.order | scripts/mod/modpost -m -a  -o /home/edinmv/go/src/Practica_2/backend/modules/moduloram/Module.symvers -e -i Module.symvers   -T -
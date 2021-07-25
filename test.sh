#!/bin/bash

make clean   >/dev/null 
make 2>&1    >/dev/null 
rm -f out.txt
./bin/Paxos_Test > out.txt
# yy=$(tail -n 2 /root/out.txt  | awk '{print $5}')
# echo ${yy:18:20}  

#!/bin/bash
MODULES="crypto net oauth2 sys text tools"
for module in ${MODULES}
do
    wget https://github.com/golang/${module}/archive/master.tar.gz -O D:/workspace/goworkspace/src/golang.org/x/${module}.tar.gz
    cd D:/workspace/goworkspace/src/golang.org/x && tar zxvf ${module}.tar.gz && mv ${module}-master/ ${module}
done

wget https://github.com/google/go-genproto/archive/master.tar.gz -O D:/workspace/goworkspace/src/google.golang.org/genproto.tar.gz
cd D:/workspace/goworkspace/src/google.golang.org && tar zxvf genproto.tar.gz && mv go-genproto-master genproto

#!/bin/zsh

BASEPATH=$(cd `dirname $0`; pwd)

echo "BASEPATH: ${BASEPATH}"

goctl api go --api ${BASEPATH}/../restful/user/user.api --dir ${BASEPATH}/../restful/user/ --home="${BASEPATH}/../templates"

#cd ./restful/user/
#goctl api doc  --dir ../../templates/
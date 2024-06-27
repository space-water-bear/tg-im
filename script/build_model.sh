#!/bin/zsh

BASEPATH=$(cd `dirname $0`; pwd)

echo "BASEPATH: ${BASEPATH}"

module="$1"
category="$2"

db_host="10.3.21.250"
db_port="5432"
db_user="postgres"
db_password="1234567890"
db_name="tg-im"

goctl model pg datasource --url="postgresql://${db_user}:${db_password}@${db_host}:${db_port}/${db_name}?sslmode=disable" \
  --table "users" \
  --home="${BASEPATH}/../templates" \
  -d "${BASEPATH}/../$category/$module/internal/model"
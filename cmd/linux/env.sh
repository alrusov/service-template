#!/bin/bash

set -a

APP_DIR=$(dirname $0)

APP_EXEC=$(readlink -f ${APP_DIR}/APP)
if [ $? != 0 ]; then
  echo no APP link file found
  exit
fi

APP=$(basename ${APP_EXEC})

LOG_LEVEL=TRACE4
PORT=1234

GODEBUG=madvdontneed=1

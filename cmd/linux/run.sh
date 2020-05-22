#!/bin/bash

source $(dirname $0)/env.sh
./${APP} --config config/${APP}.toml

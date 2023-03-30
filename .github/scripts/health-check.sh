#!/bin/bash

# if container is not running
if [ ! "$(docker ps -a -q -f name=$APP_NAME -f status=running)" ]; then
  # exit with error
  echo 1
else
  echo 0
fi
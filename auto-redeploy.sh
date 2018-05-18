#!/bin/bash

LOCK_FILE=redeploy.lock
IMAGE=docs/docker.github.io:latest
SERVICE_NAME=docs_site

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

lock() {
  date -u > $LOCK_FILE
}

is_locked() {
  if [ -e $LOCK_FILE ]; then
    return 0
  fi
  return 1
}

same_image() {
  local old=$1
  local new=$2
  if [ $old == $new ]; then
    return 0
  fi
  return 1
}

print() {
  local msg=$1
  echo -e "$BLUE[$$]$NC $msg"
}

unlock_and_exit() {
  local return_code=$1
  print finished
  rm -f $LOCK_FILE
  exit $return_code
}

main() {
  print "Redeploy script waking up: $date"
  if is_locked ; then
    print "${GREEN}Previous job still running, skipping${NC}"
    exit 0
  fi
  lock
  local curr_image="$(docker images -q $IMAGE)"
  if [ -z $curr_image ]; then
    # This should not be possible
    print "${GREEN}Image doesn't exist yet, skipping${NC}"
    unlock_and_exit 0
  fi
  docker pull $IMAGE
  local new_image="$(docker images -q $IMAGE)"
  if same_image $curr_image $new_image ; then
    print "${GREEN}No new image to deploy, skipping${NC}"
    unlock_and_exit 0

  fi
  print "Redeploying service"
  docker service update --update-parallelism 1 --image $IMAGE $SERVICE_NAME
  unlock_and_exit
}

main

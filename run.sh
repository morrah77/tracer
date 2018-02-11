#! /bin/sh
EXISTS=`docker ps -a --filter Name=tracer --format '{{.Status}}'`
if [ -n "$EXISTS" ]; then
  case $EXISTS in
    *Up*)
    echo 'Stop running container...'
    docker stop tracer;;
  esac
  echo 'Remove stopped container...'
  docker rm tracer
fi
echo 'Start new container...'
docker run -d --name=tracer tracer /tracer/bin/server -listen 8088
echo 'Started at '`docker ps --filter Name=tracer --format '{{.IP}}'`

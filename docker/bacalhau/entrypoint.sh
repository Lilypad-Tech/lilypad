#!/bin/bash

find /run /var/run -iname 'docker*.pid' -delete
dind dockerd & 
until pgrep "dockerd" >/dev/null; do sleep 0.5; done
exec "$@"

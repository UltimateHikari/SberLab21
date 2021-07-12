#!/bin/bash
apt update
apt install -y docker.io
echo ${backend_eip_id}:8000 > /root/ip.txt
docker run -d -p 80:80 -e VUE_APP_ROOT_API=${backend_eip_id}:8000 ultimatehikari/sberlab_07_21:front > /root/log.txt

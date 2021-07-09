#!/bin/bash
apt update
apt install -y docker.io
docker run -d -p 80:80 ultimatehikari/sberlab_07_21:front > /root/log.txt

#!/bin/bash
apt update
apt install -y docker.io
docker run -d -p 8000:8000 ultimatehikari/sberlab_07_21:apiserver > /root/log.txt

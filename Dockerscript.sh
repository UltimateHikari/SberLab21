docker build ./front -t ultimatehikari/sberlab_07_21:front
docker build ./apiserver -t ultimatehikari/sberlab_07_21:apiserver
docker push ultimatehikari/sberlab_07_21:front 
docker push ultimatehikari/sberlab_07_21:apiserver

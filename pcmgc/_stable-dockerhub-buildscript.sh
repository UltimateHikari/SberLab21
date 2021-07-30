if [ -z "$4" ]
  then
    echo "4 versions required, front/api/file/effect"
fi


docker build ./front -t ultimatehikari/pcmgc-front:$1
docker build ./apiserver -t ultimatehikari/pcmgc-apiserver:$2
docker build ./fileserver -t ultimatehikari/pcmgc-fileserver:$3
docker build ./effectserver -t ultimatehikari/pcmgc-effect:$4

echo "Do you wish to push builds:\
    \npcmgc-front:$1\
    \npcmgc-apiserver:$2\
    \npcmgc-fileserver:$3\
    \npcmgc-effect:$4?"
select yn in "Yes" "No"; do
    case $yn in
        Yes )  
            docker login
            docker push ultimatehikari/pcmgc-front:$1
            docker push ultimatehikari/pcmgc-apiserver:$2
            docker push ultimatehikari/pcmgc-fileserver:$3
            docker push ultimatehikari/pcmgc-effect:$4
            break;;
        No ) exit;;
    esac
done


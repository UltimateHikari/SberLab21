if [ -z "$4" ]
  then
    echo "Not enough tags"
fi
echo "Do you wish to push this builds?"
select yn in "Yes" "No"; do
    case $yn in
        Yes ) 
            echo "success"; 
            echo "success"; 
            echo "success"; 
            echo "success"; 
            break;;
        No ) exit;;
    esac
done


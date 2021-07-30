echo "Pending pods: "
kubectl get pods | grep Pending | wc -l

EXTERNAL_IP=$(kubectl -n kourier-system get service kourier -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
echo EXTERNAL_IP=$EXTERNAL_IP
KNATIVE_DOMAIN="$EXTERNAL_IP.nip.io"
echo KNATIVE_DOMAIN=$KNATIVE_DOMAIN
dig $KNATIVE_DOMAIN
kubectl patch configmap -n knative-serving config-domain -p "{\"data\": {\"$KNATIVE_DOMAIN\": \"\"}}"
kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'


export KNATIVE_VERSION="0.23.0"
kubectl apply -f https://github.com/knative/serving/releases/download/v$KNATIVE_VERSION/serving-crds.yaml
kubectl wait --for=condition=Established --all crd

kubectl apply -f https://github.com/knative/serving/releases/download/v$KNATIVE_VERSION/serving-core.yaml

kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n knative-serving > /dev/null

export KNATIVE_NET_KOURIER_VERSION="0.23.0"
kubectl apply -f https://github.com/knative/net-kourier/releases/download/v$KNATIVE_NET_KOURIER_VERSION/kourier.yaml
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n kourier-system
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n knative-serving


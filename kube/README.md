### Guide on deployment
* provide env variables stated in root readme
* cd terraform && terraform apply
* provide kublr to created cluster
* check eip in cm
* k apply -f ./volumes/cm-front-gen.yaml
* k apply -f ./services
* k apply -f ./deployments

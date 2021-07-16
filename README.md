## PiCMaGiC
### Info
Webapp for applying random transforms to images, both stored and uploaded.
Also some prerequisite subprojects for getting started with Go- cloud cli & simple server.
Educational purpose - getting familiar with Kubernetes and clouds via tinkering with services

### Features
* Vue + Bootstrap + Go
* As many services as possible
* Irrational effect application for easier max load research

### Notes
* Images on DockerHub: ultimatehikari/pcmgc-${serviceName}:${version}
* Cli expects file ../keys with:
    access-key\nsecret-key
* Terraform expects enviroment variables SBC_REGION_NAME, SBC_ACCESS_KEY, SBC_SECRET_KEY, SBC_PROJECT_NAME

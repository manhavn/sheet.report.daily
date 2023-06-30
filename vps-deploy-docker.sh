#!/bin/bash
# shellcheck disable=SC2164
cd "$(dirname "$0")"

DI_PROJECT=manhavn
DI_PACKAGE=sheetreportdaily
#DI_PACKAGE2=sheetreportdaily2
DI_VERSION=v0.0.1

if [ "$1" ]; then
  DI_VERSION=$1
fi

sh format.sh
go mod tidy
go mod vendor

docker build -t $DI_PROJECT/$DI_PACKAGE:$DI_VERSION .
docker image prune --filter="dangling=true"
docker push $DI_PROJECT/$DI_PACKAGE:$DI_VERSION

ssh root@"${DI_PACKAGE}.autoketing.org" "docker tag $DI_PROJECT/$DI_PACKAGE:$DI_VERSION $DI_PROJECT/$DI_PACKAGE"
ssh root@"${DI_PACKAGE}.autoketing.org" "docker rmi $DI_PROJECT/$DI_PACKAGE:$DI_VERSION"
ssh root@"${DI_PACKAGE}.autoketing.org" "docker pull $DI_PROJECT/$DI_PACKAGE:$DI_VERSION"
ssh root@"${DI_PACKAGE}.autoketing.org" "docker stop $DI_PACKAGE && docker rm $DI_PACKAGE"
ssh root@"${DI_PACKAGE}.autoketing.org" "docker run -d --restart always --name $DI_PACKAGE -v /srv/sqlite:/app/sqlite -p 172.17.0.1:7111:8080 -it $DI_PROJECT/$DI_PACKAGE:$DI_VERSION"
#sleep 60
#ssh root@"${DI_PACKAGE}.autoketing.org" "docker stop $DI_PACKAGE2 && docker rm $DI_PACKAGE2"
#ssh root@"${DI_PACKAGE}.autoketing.org" "docker run -d --restart always --name $DI_PACKAGE2 -v /srv/sqlite:/app/sqlite -p 172.17.0.1:7002:8080 -it $DI_PROJECT/$DI_PACKAGE:$DI_VERSION"
ssh root@"${DI_PACKAGE}.autoketing.org" "docker rmi $DI_PROJECT/$DI_PACKAGE"

# docker run -d --restart always --name sheetreportdaily -v /srv/sqlite:/app/sqlite -p 8080:8080 -it manhavn/sheetreportdaily:v0.0.1
#docker run -d --restart always -p 127.0.0.1:11211:11211 --name memcached -it memcached:alpine
#docker run -d --restart always -p 11211:11211 --name memcached -it memcached:alpine
#docker run -d --restart always --network host --name sheetreport -p 172.17.0.1:8080:8080 -it manhavn/sheetreport:v0.0.1

# scp -r * root@report.autoketing.org:/root/engine/

# /root/engine/sheet_report/up-mongo
#sh vps-deploy-docker.sh
#sh nginx-memcached/upnginx.sh
#docker start mongo-express

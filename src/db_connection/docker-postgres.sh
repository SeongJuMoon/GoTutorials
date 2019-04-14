#!/bin/bash

DB_PORT=54
CONTAINER_NAME="postgres"
CONTAINER_IMAGE=$(docker images | grep ${CONTAINER_NAME})
CONTAINER_IS_RUNNING=$(docker ps | grep ${CONTAINER_NAME})

# 이미지가 없으면 public hub에서 가져온다.
if [ ! "${CONTAINER_IMAGE}" ]; then
	echo ${CONTAINER_NAME}  "image not pulled pull \\n public registry (https://hub.docker.com)"
	docker pull postgres:latest
fi

# 실행 중인지 확인한 이후에 실행중이면 종료하고 아니면 다시 기동한다.

echo "check container is running now"

if [ "${CONTAINER_IS_RUNNING}" ]; then
	echo "container is running " ${CONTAINER_NAME} 
	docker stop ${CONTAINER_NAME}
	docker rm ${CONTAINER_NAME}
else
	echo "container run task by name " ${CONTAINER_NAME} 
	docker run -d --name ${CONTAINER_NAME} \
	-p 5432:5432 \
	-e POSTGRES_USER="web-user" \
	-e POSTGRES_PASSWORD="1q2w3e4r" \
	${CONTAINER_NAME}
fi


unset CONTAINER_NAME
unset CONTAINER_IMAGE
unset CONTAINER_IS_RUNNING
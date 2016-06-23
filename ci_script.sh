#!/usr/bin/env bash

docker ps -q -a | xargs docker rm
docker rmi $(docker images | grep “^<none>” | awk ‘{print $3}’)
docker build -t snabar/core-service .
echo "$(tput setaf 2)Building image $(tput sgr0)"
    docker stop snabar-core-service
    echo "$(tput setaf 1)Stopping Docker image $(tput sgr0)"
    docker rm snabar-core-service
    echo "$(tput setaf 1)Removing Docker image $(tput sgr0)"
docker run --name snabar-core-service -d snabar/core-service
echo "$(tput setaf 2)Running container for the first time $(tput sgr0)"

if [ $? != 0 ]
then
    docker stop snabar-core-service
    echo "$(tput setaf 1)Stopping Docker image $(tput sgr0)"
    docker rm snabar-core-service
    echo "$(tput setaf 1)Removing Docker image $(tput sgr0)"
    docker run -d --rm --name snabar-core-service -d snabar/core-service
    echo "$(tput setaf 1)Starting new Docker image $(tput sgr0)"
fi


echo "$(tput setaf 2)Deployed Coming-Soon \n Please visit http://snabar.com $(tput sgr0)"
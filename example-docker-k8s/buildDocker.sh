if [ -z $1 ]; then
    docker build -t example-docker-img .
else
    docker build -t $1 .
fi
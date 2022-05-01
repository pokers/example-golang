if [ -z $1 ]; then
    docker run -d -p 8080:8080 --name example-docker-img example-docker-img
else
    docker run -d -p 8080:8080 --name $1 $1
fi
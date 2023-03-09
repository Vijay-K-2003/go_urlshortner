### Commands

#### To remove containers and images for this project
```
docker container stop go_urlshortner-api-1 go_urlshortner-db-1
docker rm go_urlshortner-api-1 go_urlshortner-db-1
docker rmi go_urlshortner-api go_urlshortner-db
```

#### To run docker compose in detached mode
```
sudo docker compose up -d
```

#### To install required packages
```
cd api
go mod tidy
cd ..
```
# trainee-assignment  
> http://localhost/swagger/#  
> http://localhost/api/v2  

## Requirements  
- Golang 1.21;  

## Installation  
```shell
git clone https://github.com/go-swagger/go-swagger
cd go-swagger
go install ./cmd/swagger
```

## Running

### Dev  
```shell
make up  
```
or
```shell
docker-compose -f docker-compose.dev.yml up --build -d  
```

### Tests  
```shell
make tests  
```
or
```shell
docker-compose -f docker-compose.test.yml up --abort-on-container-exit --exit-code-from tests --quiet-pull
```

### Common
```shell
docker-compose -f docker-compose.test.yml up --build  
```

## Swagger generation  
```shell
make generate  
```

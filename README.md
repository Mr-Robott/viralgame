# viralgame

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

ViralGame is a simple service that handles the User Data.
## Pre - requisite

- Go (1.17)
- Docker
- Docker Compose
- MySQL Server

## Installation

- If you want to run directly using the docker compose, go to the root directory of project and run below command. Make sure you export all the envs in .env

```sh
docker-compose up --build
```

## Routes
- ```/user``` [POST] => *Create User*

From other terminal. Do a curl request.

```sh
curl -d '{"name": "Pallav"}' -X POST http://127.0.0.1:5000/user
```

The output would be something like this.

```sh
{
  "id": "40f41829-dfab-4f5b-ae49-4e90658196f2",
  "name": "Pallav"
}
``` 


- ```/user/<id>/state``` [PUT] => *Update Game State*

Request -
```sh
curl -d '{"gamesPlayed": 422, "score": 6000}' -X PUT http://127.0.0.1:5000/user/40f41829-dfab-4f5b-ae49-4e90658196f2/state
```

Response -
```sh
{
  "user": {
    "id": "40f41829-dfab-4f5b-ae49-4e90658196f2",
    "name": "Pallav"
  },
  "gameState": {
    "gamesPlayed": 422,
    "score": 6000
  },
  "created_at": "2022-05-16T22:49:06Z",
  "updated_at": "2022-05-16T23:10:01Z"
}
```


- ```/user/<id>/state``` [GET] => *Load Game State*

Request -
```sh
curl http://127.0.0.1:5000/user/40f41829-dfab-4f5b-ae49-4e90658196f2/state
```

Response -
```sh
{
  "gamesPlayed": 422,
  "score": 6000
}
```
- ```/user/<id>/friends``` [PUT] => *Update Friends*
  Request -
```sh
curl -d '{"friends": ["7eeb1f9a-8f6d-4769-a670-a685f4add450","8f0dca46-e7b7-432a-864e-5c3e5533e14a"]}' -X PUT http://127.0.0.1:5000/user/40f41829-dfab-4f5b-ae49-4e90658196f2/friends
```
- ```/user/<id>/friends``` [PUT] => *Get All Friends*
- Request -
```sh
curl http://127.0.0.1:5000/user/40f41829-dfab-4f5b-ae49-4e90658196f2/friends 
```
- Response -
```sh
[
  {
    "id": "7eeb1f9a-8f6d-4769-a670-a685f4add450",
    "named": "Hector",
    "highScore": 2132130
  },
  {
    "id": "8f0dca46-e7b7-432a-864e-5c3e5533e14a",
    "named": "MrRobot",
    "highScore": 786567
  }
]
```
- ```/user``` [GET] => *Get All Users*


Response -

```sh
{
  "users": [
    {
      "id": "295c6a9d-017d-4a31-8700-fa5dac18dd92",
      "name": "Dave"
    },
    {
      "id": "2a41f290-6d09-4423-9ba2-ae63d9de0948",
      "name": "Pallav"
    },
    {
      "id": "40b4007d-68a1-449e-84f5-4e9d12cdd23b",
      "name": "Pallav"
    },
    {
      "id": "40f41829-dfab-4f5b-ae49-4e90658196f2",
      "name": "Pallav"
    },
    {
      "id": "7eeb1f9a-8f6d-4769-a670-a685f4add450",
      "name": "Hector"
    },
    {
      "id": "8f0dca46-e7b7-432a-864e-5c3e5533e14a",
      "name": "MrRobot"
    }
  ]
}
```

```If you want to manually run the service, go to the root directory of project and run. Make sure of you have free port at 5000 and have mysql-server up and running and defined values for mysql db are correct in .env```


```sh
go run main.go
```
- If you want to manu

or you can build binary and run the same.

## License

Pallav Mathur


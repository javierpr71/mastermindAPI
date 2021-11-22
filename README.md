# MasterMind

## Prerequisites

- Redis instance

## Configuration

Configure the next environment variables:
 
| Var | Description |
|-----|-------------|
| PORT| Mastermind Api will start in this port|
| REDIS| Url of redis instance in the form <host>:<port>. I.e: localhost:6379|


## Build

A Go compiler and make tool need to be installed.

In a command shell run:

```
make build-all
```

The result is a bin folder with binary file for macos, linux and windows.

## Running

After build, go to bin folder and choose the binary for you apropiete architecture

In a command shell run the selected binary.


## Deploy

- Configure environment variables
- Deploy apropiate binary located in bin folder

## MasteMind API Endpoints

For testing purpouses, a Postman file is included in forlder ./tests/postman 

```
POST http://localhost:8080/newgame
```

### Responses
| Name   | Type   | Description          |
|--------|--------|----------------------|
| 200 Ok | string | Ok. UUID of new game |


```
POST http://localhost:8080/round
```

## Request Body

|Name   |Type    |Description            |
|-------|--------|-----------------------|
| id    | string | UUID Id of the game   |
| guess | string | String to evaluate    |

```
GET http://localhost:8080/status/{id}
```
## URI Parameters

| Name | in   | Required | Type        | Description |
|------|------|----------|-------------|-------------|
| id   | Path | true     | string uuid | Game Id     |


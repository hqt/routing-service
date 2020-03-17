# Routing Service
You are provided data on the stations and lines of Singapore's urban rail system, including planned additions over the next few years. Your task is to use this data to build a routing service, to help users find routes from any station to any other station on this future network.

## Technical Detail
- Language: Golang 1.14
- Other: Docker
- Simple routing API: BFS algorithm
- Advanced routing API: Dijkstra algorithm
- Application will be hosted on port 3000

## API
This service has 2 endpoints.

### `POST api/simple_route`
Parameters:

|Name|Description|Required?|
|----|-----------|---------|
|start|Start station|Yes|
|stop|Destination station|Yes|


### `POST api/advanced_route`
Parameters:

|Name|Description|Required?|
|----|-----------|---------|
|start|Start station|Yes|
|stop|Destination station|Yes|
|time|Start time ("YYYY-MM-DDThh:mm" format, e.g. '2019-01-31T16:00')|Yes|

### Example response
```json
{
    "verdict": "success",
    "minutes": 134,
    "path": [
        "EW27",
        "EW26",
        "EW25",
        "EW24",
        "EW23",
        "EW22",
        "EW21",
        "CC22",
        "CC21",
        "CC20",
        "CC19",
        "DT9",
        "DT10",
        "DT11",
        "DT12"
    ],
    "instructions": [
        "Take EW line from Boon Lay to Lakeside",
        "Take EW line from Lakeside to Chinese Garden",
        "Take EW line from Chinese Garden to Jurong East",
        "Take EW line from Jurong East to Clementi",
        "Take EW line from Clementi to Dover",
        "Take EW line from Dover to Buona Vista",
        "Change from EW line to CC line",
        "Take CC line from Buona Vista to Holland Village",
        "Take CC line from Holland Village to Farrer Road",
        "Take CC line from Farrer Road to Botanic Gardens",
        "Change from CC line to DT line",
        "Take DT line from Botanic Gardens to Stevens",
        "Take DT line from Stevens to Newton",
        "Take DT line from Newton to Little India"
    ],
    "startAt": "2019-01-31T16:00"
}
```

## Packages
This project is a set of separate of concerns, independent components, assembled
into a HTTP server to solve the problem. The components include:

- Data Structure: Implements some data structures (FIFO queue, Priority Queue) for solving problems.
- Routing: contains all logics to find routing between points
    - parser: read from csv file and parsing to the graph
    - implement BFS/Dijkstra for finding routes.
- Server: for serving http requests. Using golang native `net/http package`

## Running the project
Prerequisite: Docker for running tests and server

- Run tests
```bash
make ci-test
```

- Run server (using Docker):
```bash
make start-server
```

- Send a request to simple routing API (after starting the server)
```bash
make simple-request
```

- Send a request to advanced routing API (after starting the server)
```bash
make advanced-request
```

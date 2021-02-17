# Random numbers standard deviation

## Description

Simple rest api returns standard deviations for random numbers set and numbers set composed of sum achieved sets.

Call http GET to route below : 

```http://localhost:8080/random/mean?requests={numOfRequests}&length={length}```

Filter params reference :

``{numOfRequests}`` How many numbers set we want to achieve
``{length}`` How many numbers per set should be drawn

## Development

To build development docker image application :

``docker build -t standard-deviation-dev . -f ./docker/dev/Dockerfile``

You need obtain random api key from https://www.random.org/
To run development application provide proper .env.local basing on .env file

``docker run --name standard-deviation-dev -it --rm -p 8080:8080 standard-deviation-dev --env-file .env.local`` 

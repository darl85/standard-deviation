# Random numbers standard deviation

To build development docker image application :

``docker build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) -t standard-deviation-dev .``

To run development application :

``docker run --name standard-deviation-dev -it --rm -p 8010:8010 -v $PWD:/go/standard-deviation standard-deviation-dev``

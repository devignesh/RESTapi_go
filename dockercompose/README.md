Go with Docker 

------------------------------------------------------

1. clone the github repo

2. Build image 

    docker-compose build

3. Run the Docker-compose file

    docker-compose up


Install Docker-compose 

    1. sudo curl -L "https://github.com/docker/compose/releases/download/{docker-compose version}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

    2. sudo ln -sf /usr/local/bin/docker-compose /usr/bin/docker-compose

    3. sudo chmod +x /usr/bin/docker-compose

    4. docker-compose --version

make sure the you have docker and docker-compose in your machine 

    To verify 

    1. docker --version

    2. docker-compose --version
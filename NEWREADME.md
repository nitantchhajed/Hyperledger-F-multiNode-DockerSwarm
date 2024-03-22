commands
> docker network create hlf-network
> cd artifacts/channel
> ./create-artifacts.sh

>  sudo chmod 777 -R *
>  cd ..
>  docker-compose up -d

All the services will be up and running

>  cd ..
>  ./create-channel.sh

the channel will be created, joined, and update the anchor peer

> docker exec -it cli bash
>  peer channel list

This will list your channel
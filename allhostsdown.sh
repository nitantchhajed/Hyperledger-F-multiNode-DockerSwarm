docker-compose -f host1.yaml down -v 
docker-compose -f host2.yaml down -v 
docker-compose -f host3.yaml down -v 
docker-compose -f host4.yaml down -v 
docker-compose -f host5.yaml down -v 

docker rm $(docker ps -aq)
docker rmi $(docker images dev-* -q)
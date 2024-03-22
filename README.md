# Create Docker network

*For local -*
``docker network create hlf-network``


*For Swarm deployment- *
``docker swarm init --advertise-addr <host-1 ip address>``

``docker swarm join-token manager``
*from the output of above command* 

*on different servers-*
``<output from join-token manager> --advertise-addr <VM n ip>``

*on 1st server -*
``docker network create --attachable --driver overlay hlf-network``



*Crypto Material*

``./bin/cryptogen generate --config=./crypto-config.yaml --output="crypto-config"``

*Channel Artifacts -*

``./bin/configtxgen -profile SampleMultiNodeEtcdRaft -outputBlock ./channel-artifacts/genesis.block``
``./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel``
``./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP``


channel create-

``docker exec cli peer channel create -o orderer0.subskribo.be:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem``

# Join peers to channel-

*For peer 1 -*

``docker exec cli peer channel join -b ./channel-artifacts/genesis.block``

*For peer 2- *

``docker exec -e CORE_PEER_ADDRESS=peer1.org1.subskibo.be:7051 -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/ca.crt cli peer channel join -b ./channel-artifacts/genesis.block``

# Chaincode Installation -*

make sure to replace the Path with your chaincode path

*For peer 1-*

``docker exec cli peer chaincode install -n mycc -v 1.0 -p ./chaincode``

*For peer 2-*

``docker exec -e CORE_PEER_ADDRESS=peer1.org1.subskribo.be:7051 -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/ca.crt cli peer chaincode install -n mycc -v 1.0 -p ./chaincode``


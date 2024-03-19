
``../fabric-samples/bin/configtxgen -profile SampleMultiNodeEtcdRaft -outputBlock ./channel-artifacts/genesis.block``
``../fabric-samples/bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel``
``../fabric-samples/bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP``


channel create-

``docker exec cli peer channel create -o orderer0.subskribo.be:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem``
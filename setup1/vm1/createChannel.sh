export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/../vm4/crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem
export PEER0_ORG1_CA=${PWD}/crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/ca.crt
export FABRIC_CFG_PATH=${PWD}/../../artifacts/channel/config/

export CHANNEL_NAME=mychannel

setGlobalsForPeer0Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

setGlobalsForPeer1Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be/msp
    export CORE_PEER_ADDRESS=localhost:8051
}

createChannel(){
    rm -rf ./channel-artifacts/*
    setGlobalsForPeer0Org1
    
    # Replace localhost with your orderer's vm IP address
    peer channel create -o 10.1.0.13:7050 -c $CHANNEL_NAME \
    --ordererTLSHostnameOverride orderer0.subskribo.be \
    -f ./../../artifacts/channel/${CHANNEL_NAME}.tx --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block \
    --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
}

#createChannel

joinChannel(){
    setGlobalsForPeer0Org1
    peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
    
}

#joinChannel

updateAnchorPeers(){
    setGlobalsForPeer0Org1
    # Replace localhost with your orderer's vm IP address
    peer channel update -o 10.1.0.13:7050 --ordererTLSHostnameOverride orderer0.subskribo.be -c $CHANNEL_NAME -f ./../../artifacts/channel/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
    
}

updateAnchorPeers

# removeOldCrypto

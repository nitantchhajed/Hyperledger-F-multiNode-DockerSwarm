createcertificatesForOrg1() {
  echo
  echo "Enroll the CA admin"
  echo
  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/
  export FABRIC_CA_CLIENT_HOME=${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/

  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca.org1.subskribo.be --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1-subskribo-be.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1-subskribo-be.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1-subskribo-be.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1-subskribo-be.pem
    OrganizationalUnitIdentifier: orderer' >${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/config.yaml

  echo
  echo "Register peer0"
  echo
  fabric-ca-client register --caname ca.org1.subskribo.be --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  echo
  echo "Register peer1"
  echo
  fabric-ca-client register --caname ca.org1.subskribo.be --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  echo
  echo "Register user"
  echo
  fabric-ca-client register --caname ca.org1.subskribo.be --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  echo
  echo "Register the org admin"
  echo
  fabric-ca-client register --caname ca.org1.subskribo.be --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/peers

  # -----------------------------------------------------------------------------------
  #  Peer 0
  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be

  echo
  echo "## Generate the peer0 msp"
  echo
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/msp --csr.hosts peer0.org1.subskribo.be --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/config.yaml ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls --enrollment.profile tls --csr.hosts peer0.org1.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/ca.crt
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/signcerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/server.crt
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/keystore/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/server.key

  mkdir ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/tlscacerts
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/tlscacerts/ca.crt

  mkdir ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/tlsca
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/tlsca/tlsca.org1.subskribo.be-cert.pem

  mkdir ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/ca
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer0.org1.subskribo.be/msp/cacerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/ca/ca.org1.subskribo.be-cert.pem

  # ------------------------------------------------------------------------------------------------

  # Peer1

  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be

  echo
  echo "## Generate the peer1 msp"
  echo
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/msp --csr.hosts peer1.org1.subskribo.be --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/config.yaml ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/msp/config.yaml

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls --enrollment.profile tls --csr.hosts peer1.org1.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/ca.crt
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/signcerts/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/server.crt
  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/keystore/* ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/peers/peer1.org1.subskribo.be/tls/server.key

  # --------------------------------------------------------------------------------------------------

  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/users
  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/users/User1@org1.subskribo.be

  echo
  echo "## Generate the user msp"
  echo
  fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/users/User1@org1.subskribo.be/msp --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  mkdir -p ../crypto-config/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be

  echo
  echo "## Generate the org admin msp"
  echo
  fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca.org1.subskribo.be -M ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be/msp --tls.certfiles ${PWD}/fabric-ca/org1/tls-cert.pem

  cp ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/msp/config.yaml ${PWD}/../crypto-config/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be/msp/config.yaml

}

createcertificatesForOrg1

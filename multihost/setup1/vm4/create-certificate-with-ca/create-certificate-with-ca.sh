createCretificateForOrderer() {
  echo
  echo "Enroll the CA admin"
  echo
  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be

  export FABRIC_CA_CLIENT_HOME=${PWD}/../crypto-config/ordererOrganizations/subskribo.be

  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' >${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/config.yaml

  echo
  echo "Register orderer0"
  echo

  fabric-ca-client register --caname ca-orderer --id.name orderer0 --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  echo
  echo "Register orderer1"
  echo

  fabric-ca-client register --caname ca-orderer --id.name orderer1 --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  echo
  echo "Register orderer2"
  echo

  fabric-ca-client register --caname ca-orderer --id.name orderer2 --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  echo
  echo "Register the orderer admin"
  echo

  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/orderers
  # mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/orderers/subskribo.be

  # ---------------------------------------------------------------------------
  #  Orderer0

  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be

  echo
  echo "## Generate the orderer0 msp"
  echo

  fabric-ca-client enroll -u https://orderer0:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp --csr.hosts orderer0.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/config.yaml ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/config.yaml

  echo
  echo "## Generate the orderer0-tls certificates"
  echo

  fabric-ca-client enroll -u https://orderer0:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls --enrollment.profile tls --csr.hosts orderer0.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/ca.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/signcerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/server.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/keystore/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/server.key

  mkdir ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/tlscacerts
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem

  mkdir ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/tlscacerts
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer0.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem

  # -----------------------------------------------------------------------
  #  Orderer1

  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be

  echo
  echo "## Generate the orderer1 msp"
  echo

  fabric-ca-client enroll -u https://orderer1:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/msp --csr.hosts orderer1.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/config.yaml ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/msp/config.yaml

  echo
  echo "## Generate the orderer1-tls certificates"
  echo

  fabric-ca-client enroll -u https://orderer1:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls --enrollment.profile tls --csr.hosts orderer1.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/ca.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/signcerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/server.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/keystore/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/server.key

  mkdir ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/msp/tlscacerts
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer1.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem

  # ---------------------------------------------------------------------------
  #  Orderer2
  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be

  echo
  echo "## Generate the orderer2 msp"
  echo

  fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/msp --csr.hosts orderer2.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/config.yaml ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/msp/config.yaml

  echo
  echo "## Generate the orderer2-tls certificates"
  echo

  fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls --enrollment.profile tls --csr.hosts orderer2.subskribo.be --csr.hosts localhost --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/ca.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/signcerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/server.crt
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/keystore/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/server.key

  mkdir ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/msp/tlscacerts
  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/tls/tlscacerts/* ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/orderers/orderer2.subskribo.be/msp/tlscacerts/tlsca.subskribo.be-cert.pem
  # ---------------------------------------------------------------------------

  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/users
  mkdir -p ../crypto-config/ordererOrganizations/subskribo.be/users/Admin@subskribo.be

  echo
  echo "## Generate the admin msp"
  echo

  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/users/Admin@subskribo.be/msp --tls.certfiles ${PWD}/fabric-ca/ordererOrg/tls-cert.pem

  cp ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/msp/config.yaml ${PWD}/../crypto-config/ordererOrganizations/subskribo.be/users/Admin@subskribo.be/msp/config.yaml

}

createCretificateForOrderer
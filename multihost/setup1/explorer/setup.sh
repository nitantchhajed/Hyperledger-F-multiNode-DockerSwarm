export ORG1_MSPKEY=$(cd ../test-network/organizations/peerOrganizations/org1.subskribo.be/users/Admin@org1.subskribo.be/msp/keystore && ls *_sk) 


sed -i  "s/ORG1_MSPKEY/$ORG1_MSPKEY/g" test.json 
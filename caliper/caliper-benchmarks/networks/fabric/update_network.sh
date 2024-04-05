#!/bin/bash

# Change permissions for the first-network directory and its subdirectories
chmod -R a+rw "${PWD}/first-network"

# Find the priv_sk and User1@org1.example.com-cert.pem files in the directory
keystore_file_path_org1=$(find "${PWD}/first-network/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore" -type f)
cert_file_path_org1=$(find "${PWD}/first-network/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts" -type f)
keystore_file_org1=$(basename "$keystore_file_path_org1")
cert_file_org1=$(basename "$cert_file_path_org1")

# Find the priv_sk and User1@org2.example.com.pem files in the directory
keystore_file_path_org2=$(find "${PWD}/first-network/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore" -type f)
cert_file_path_org2=$(find "${PWD}/first-network/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/signcerts" -type f)
keystore_file_org2=$(basename "$keystore_file_path_org2")
cert_file_org2=$(basename "$cert_file_path_org2")

if [ -n "$keystore_file_path_org1" ] && [ -n "$cert_file_org1" ] && [ -n "$keystore_file_path_org2" ] && [ -n "$cert_file_org2" ]; then
  echo "Found keystore file for org1: $keystore_file_org1"
  echo "Found cert file for org1: $cert_file_org1"
  echo "Found keystore file for org2: $keystore_file_org2"
  echo "Found cert file for org2: $cert_file_org2"
  # Replace the "path" value with keystore_file
  config_json='name: Caliper Benchmarks
version: "2.0.0"

caliper:
  blockchain: fabric

info:
  Version: 2.2.0
  Size: 2 Org with 2 Peers
  Orderer: HCS
  Distribution: Single Host
  StateDB: golevelDB

channels:
  # channelName of mychannel matches the name of the channel created by test network
  - channelName: mychannel
    # the chaincodeIDs of all the fabric chaincodes in caliper-benchmarks
    contracts:
    - id: mycc

organizations:
  - mspid: Org1MSP
    identities:
      certificates:
      - name: '"'User1'"'
        clientPrivateKey:
          path: '"'"../first-network/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/$keystore_file_org1"'"'
        clientSignedCert:
          path: '"'../first-network/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/$cert_file_org1'"'
    connectionProfile:
      path: '"'../first-network/connection-org1.yaml'"'
      discover: true
  - mspid: Org2MSP
    identities:
      certificates:
      - name: '"'User1'"'
        clientPrivateKey:
          path: '"'"../first-network/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore/$keystore_file_org2"'"'
        clientSignedCert:
          path: '"'../first-network/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/signcerts/$cert_file_org2'"'
    connectionProfile:
      path: '"'../first-network/connection-org2.yaml'"'
      discover: true'
      
  echo "$config_json" > "${PWD}/caliper/caliper-benchmarks/networks/fabric/test-network.yaml"

  echo "Configuration file updated."
else
  echo "Error: keystore file or cert file not found."
fi

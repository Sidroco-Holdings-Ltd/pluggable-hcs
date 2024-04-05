make configtxgen configtxlator cryptogen orderer peer docker

cd first-network

 ./byfn.sh up -t 20

 cd ..

 make caliper-test-up

# DeployingAndCallingContract-EVM

This repo is used for deploying and calling solidity contract for evm compatible blockchain networks

This is only for testing it was not tested for production

Enviroment i used:

Network: PolygonEdge https://github.com/0xPolygon/polygon-edge

## How to use it

1. Install solidity compiler and abigen tool
2. Add network and account to deploy.go and callMethod.go
  ```
  client, err := ethclient.Dial("<NETWORK>") //example http://localhost:10002
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("<PRIVATE KEY>") //account private key
	if err != nil {
		log.Fatal(err)
	}
  ```
3. go run deploy.go
  ```
  output:
  Contract address 0x91f273AB0d9cf7cD03D351a3E98008AfDB220f9E
  Tx hash 0xc96eae2a235b6932fd04c67a2393691d9cf6bcc0dba0868384e7613469eb44ee
  ```
4. Copy contract address to callMethod.go
  ```
  address := common.HexToAddress("<CONTRACT ADDRESS>")
  ```
5. go run callMethod.go
  ```
  Output:
  tx sent:  0xb34efe3cca804ef71395a323b7c7ffef9814e2a54ed986e45011ea361a17998d
  ```
  
## How to test it

You can test by subscribing to the logs, https://github.com/Aleksao998/SubscribeToEvents-EVM




# Crypto Service

## Description
This is a small common crypto librarly to encrypt/decrypt secret data. You can decrypt data that's encrypted through [ Crypto Service CSharp ](https://github.com/itera-io/crypto-service-csharp) and vice versa

## Installation
Use ` go get github.com/itera-io/crypto-service-go ` to install TaikunCryptoGoClient

## Usage
After installation you can create crypto client by dependency injection `client, _ := crypto.NewClient()` 
Note: Please specify `CRYPTO_SERVICE_KEY` environment variable

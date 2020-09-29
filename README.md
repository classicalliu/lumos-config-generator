# lumos config generator

A tool to generate basic config for lumos. Include `SECP256K1_BLAKE160`, `SECP256K1_BLAKE160_MULTISIG`, `DAO` scripts.

## Usage

```bash
go build

# generate config file
# file-path default to ./config.json
# rpc-url default to http://127.0.0.1:8114
./lumos-config-generator
# OR
./lumos-config-generator file-path rpc-url
# Such as
./lumos-config-generator config.json http://127.0.0.1:8114
```

## Example

```json
{
  "PREFIX": "ckt",
  "SCRIPTS": {
    "SECP256K1_BLAKE160": {
      "CODE_HASH": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
      "HASH_TYPE": "type",
      "TX_HASH": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37",
      "INDEX": "0x0",
      "DEP_TYPE": "dep_group",
      "SHORT_ID": 0
    },
    "SECP256K1_BLAKE160_MULTISIG": {
      "CODE_HASH": "0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8",
      "HASH_TYPE": "type",
      "TX_HASH": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37",
      "INDEX": "0x1",
      "DEP_TYPE": "dep_group",
      "SHORT_ID": 1
    },
    "DAO": {
      "CODE_HASH": "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e",
      "HASH_TYPE": "type",
      "TX_HASH": "0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f",
      "INDEX": "0x2",
      "DEP_TYPE": "code"
    }
  }
}
```

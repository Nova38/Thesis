

## Groups

All will be run with:

(Collections are special and can only be created by an admin user)

- no authentication
- Identity based Authentication
- Role based Authentication with no parents
- Role based Authentication with parents


### Types of Benchmarks

### Parameters:

- Number of item types
- Number of items per type
- Per Item Type:
  - Number of Items
  - Range of the number of Sub Items
  - Percentage of users who don't have permission for their operations
  - Percentage of operations that are:
    - reads
    - creates
    - updates
    - deletes 
    - suggestions (subitem with multiple values)
    - hidden (subitem with single value)

#### Operation Types
- Create and Read
- Create, Read, Update, Delete
- With Suggestions
- With Hidden 
- All

--security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE


docker run --rm -d --name peer0org1_test_ccaas --network fabric_test -p 40000:40000 -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999 -e CHAINCODE_ID=test_1.0.1:2e8b83f5fec1cb15ee64a09f690e1ca30518ae3c71b620dd00d384ea1a447593 -e CORE_CHAINCODE_ID_NAME=test_1.0.1:2e8b83f5fec1cb15ee64a09f690e1ca30518ae3c71b620dd00d384ea1a447593 test_ccaas_image:latest 


docker run --rm -d --name peer0org2_test_ccaas --network fabric_test -e CHAINCODE_SERVER_ADDRESS=0.0.0.0:9999 -e CHAINCODE_ID=test_1.0.1:2e8b83f5fec1cb15ee64a09f690e1ca30518ae3c71b620dd00d384ea1a447593 -e CORE_CHAINCODE_ID_NAME=test_1.0.1:2e8b83f5fec1cb15ee64a09f690e1ca30518ae3c71b620dd00d384ea1a447593 test_ccaas_image:latest --security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE


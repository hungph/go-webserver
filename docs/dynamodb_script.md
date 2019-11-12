## Run DynamoDB Local on Docker
```shell script
$ docker run -d --name dynamodb -p 8000:8000 -v /Users/hungphanhuu/Documents/Working/Personal/201911/dynamodb-data:/data/ amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath /data
```
## Create Table for running
```shell script
$ /usr/local/bin/aws dynamodb create-table \
    --table-name user \
    --attribute-definitions \
        AttributeName=UserNm,AttributeType=S \
    --key-schema AttributeName=UserNm,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=10 \
    --endpoint-url http://localhost:8000
```
## Create Table for testing
```shell script
$ /usr/local/bin/aws dynamodb create-table \
      --table-name user_test \
      --attribute-definitions \
          AttributeName=UserNm,AttributeType=S \
      --key-schema AttributeName=UserNm,KeyType=HASH \
      --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=10 \
      --endpoint-url http://localhost:8000
```
## Check created tables
```shell script
$ /usr/local/bin/aws dynamodb list-tables \
      --endpoint-url http://localhost:8000
```




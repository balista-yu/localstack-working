#!/bin/bash

# create s3
awslocal s3 mb s3://working-demo-bucket

# add bucket name to parameter store
#awslocal ssm put-parameter --name /working-demo/images --type "String" value "working-demo-bucket"

# create sqs
awslocal sqs create-queue --queue-name working-demo-fifo-queue.fifo --attribute "FifoQueue=true"
awslocal sqs create-queue --queue-name working-demo-queue

# create sns
awslocal sns create-topic --name working-demo-topic

# register subscription
awslocal sns subscribe \
    --topic-arn arn:aws:sns:ap-northeast-1:000000000000:working-demo-topic \
    --protocol email \
    --notification-endpoint dummy@example.com

cd /home/localstack/data/working-demo/
awslocal s3 cp . s3://working-demo-bucket --recursive
aws --endpoint-url=http://localhost:4566 s3api put-bucket-notification-configuration --bucket working-demo-bucket --notification-configuration file://event-notification.json
# event通知を設定するとTestEventのメッセージが自動で送信される仕様のため、そちらのメッセージを削除する
awslocal sqs purge-queue --queue-url "http://localhost:4566/000000000000/working-demo-queue"

# create secret-manager
awslocal secretsmanager create-secret --name local-secret-id --secret-string '{"username":"dummy", "password":"dummy"}'

# create dynamodb
awslocal dynamodb create-table \
    --table-name working-demo-table \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --billing-mode=PAY_PER_REQUEST \
    --no-deletion-protection-enabled

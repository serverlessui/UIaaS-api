```shell
$ aws cloudformation package --template-file sam.yml --output-template-file cfn-transformed-template.yml --s3-bucket lambda.bucket.moodle aws cloudformation deploy --template-file ./cfn-transformed-template.yml --stack-name lambda-sqs-processor --capabilities CAPABILITY_IAM --parameter-overrides QueueName=SomeQueueName
```

YOUR_SQS_QUEUE_URL=https://sqs.us-east-1.amazonaws.com/123456789012/my-queue; 
```shell
aws sqs send-message --queue-url $YOUR_SQS_QUEUE_URL --message-body '{ "myMessage": "Hello SAM!" }'
```
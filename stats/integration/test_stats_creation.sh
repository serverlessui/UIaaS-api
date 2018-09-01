#!/bin/bash

set -e

QUEUE_NAME=$1

QUEUE_URL=`aws sqs get-queue-url --queue-name=${QUEUE_NAME} --output text`

aws sqs send-message --queue-url $QUEUE_URL --message-body '{ "siteId": "20384s9d0fasd", "siteName": "uiaas.vssdevelopment.com" }'

 aws athena start-query-execution --query-string='SELECT count(distinct(requestip)) AS origin FROM cloudfront_logs GROUP BY  requestip;' --result-configuration=allstatebucket

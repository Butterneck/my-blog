#!/bin/bash

STREAM_ARN="arn:aws:dynamodb:eu-west-1:794148823865:table/Posts/stream/2023-11-26T11:03:27.596"
PROFILE="blog-admin"

while true
do
  # Get a list of all shards in the stream
  shards=$(aws dynamodbstreams list-shards \
    --stream-arn $STREAM_ARN \
    --query 'Shards[*].ShardId' \
    --output text \
    --profile $PROFILE)

  # Iterate through each shard and get records
  for shard_id in $shards
  do
    shard_iterator=$(aws dynamodbstreams get-shard-iterator \
      --stream-arn $STREAM_ARN \
      --shard-id $shard_id \
      --shard-iterator-type LATEST \
      --query 'ShardIterator' \
      --output text \
      --profile $PROFILE)

    records=$(aws dynamodbstreams get-records \
      --shard-iterator $shard_iterator \
      --limit 10 \
      --query 'Records[*].[Dynamodb.NewImage, SequenceNumber]' \
      --output table \
      --profile $PROFILE)

    echo "Shard ID: $shard_id, Records: $records"
  done

  sleep 1
done

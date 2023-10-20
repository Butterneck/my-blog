#!/bin/bash

export AWS_ACCESS_KEY_ID=asd
export AWS_SECRET_ACCESS_KEY=asd
export AWS_REGION=us-east-1


# Define table name
table_name="Posts"

# Point to local DynamoDB endpoint
endpoint_url="http://localhost:8000"

# Specify a region (it's optional for local DynamoDB, as it doesn't have region-specific settings)
region="us-east-1"

# Create a DynamoDB table
aws dynamodb create-table \
    --table-name $table_name \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url $endpoint_url \
    --region $region

# Check the status of the table creation (optional)
while true; do
    status=$(aws dynamodb describe-table --table-name $table_name --query 'Table.TableStatus' --output text --endpoint-url $endpoint_url --region $region)
    if [ "$status" == "ACTIVE" ]; then
        echo "Table $table_name is active."
        break
    else
        echo "Table $table_name is still being created. Status: $status"
        sleep 5
    fi
done

# Create a Global Secondary Index
aws dynamodb update-table \
    --table-name $table_name \
    --attribute-definitions AttributeName=pk,AttributeType=S AttributeName=createdAt,AttributeType=N \
    --global-secondary-index-updates \
    "[{\"Create\":{\"IndexName\":\"${table_name}-list-index\",\"KeySchema\":[{\"AttributeName\":\"pk\",\"KeyType\":\"HASH\"},{\"AttributeName\":\"createdAt\",\"KeyType\":\"RANGE\"}],\"Projection\":{\"ProjectionType\":\"ALL\"},\"ProvisionedThroughput\":{\"ReadCapacityUnits\":5,\"WriteCapacityUnits\":5}}}]" \
    --endpoint-url $endpoint_url \
    --region $region

aws dynamodb update-table \
    --table-name $table_name \
    --attribute-definitions AttributeName=slug,AttributeType=S \
    --global-secondary-index-updates \
    "[{\"Create\":{\"IndexName\":\"${table_name}-slug-index\",\"KeySchema\":[{\"AttributeName\":\"slug\",\"KeyType\":\"HASH\"}],\"Projection\":{\"ProjectionType\":\"ALL\"},\"ProvisionedThroughput\":{\"ReadCapacityUnits\":5,\"WriteCapacityUnits\":5}}}]" \
    --endpoint-url $endpoint_url \
    --region $region

echo "DynamoDB table $table_name created successfully with global secondary indexes on $endpoint_url and region $region."

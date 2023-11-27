Creates a serverless API which includes
- lambda function with container image runtime
- ecr repository to store images consumed by lambda function
- api gateway
- cognito user pool authentication
- dynamodb table with dynamodb streams enabled
- (optional) forward dynamodb changes to custom event bridge bus


To define API routes, methods, authentication and authorization an openapi 3.0 file needs to be provided.
This openapi specification file can contain references to `cognito_user_pool_arn` and `integration_uri`; these fields will be replaced by the module to the correct values.
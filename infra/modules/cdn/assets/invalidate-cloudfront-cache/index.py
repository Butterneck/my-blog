import boto3
from time import time
from os import environ

cloudfront_client = boto3.client('cloudfront')
cloudfront_distribution_id = environ['CLOUDFRONT_DISTRIBUTION_ID']


def handler(event, context):
    invalidate_paths = event['invalidate_paths']
    print(f"Invalidating CloudFront cache for {invalidate_paths}")
    resp = invalidate_cloudfront_cache(cloudfront_distribution_id, invalidate_paths)
    print(resp)
    return {}

def invalidate_cloudfront_cache(distribution_id: str, paths: list):
    response = cloudfront_client.create_invalidation(
        DistributionId=distribution_id,
        InvalidationBatch={
            'Paths': {
                'Quantity': len(paths),
                'Items': paths
            },
            'CallerReference': str(time())
        }
    )
    return response
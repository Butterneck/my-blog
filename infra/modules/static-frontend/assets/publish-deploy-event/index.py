import boto3
from os import environ

event_bus_name = environ['EVENTBRIDGE_BUS_NAME']
event_source_name = environ['EVENT_SOURCE_NAME']

def handler(event, context):
    resp = publish_event(event)
    print(resp)
    return {}

def publish_event(event):
    client = boto3.client('events')
    response = client.put_events(
        Entries=[
            {
                'Source': event_source_name,
                'DetailType': 'deploy',
                'Detail': '{}',
                'EventBusName': event_bus_name
            }
        ]
    )
    return response
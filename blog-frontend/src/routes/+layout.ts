import { Amplify } from 'aws-amplify';
import { OpenAPI } from '$lib/generated/backend-client';
import { getIdentityToken } from '$lib/auth';
import { getConfig } from '$lib/configuration/config';

const config = getConfig();

Amplify.configure({
    Auth: config.cognito
});

OpenAPI.TOKEN = async () => getIdentityToken();
// OpenAPI.BASE = "https://5znf2q3t6l.execute-api.eu-west-1.amazonaws.com/main"
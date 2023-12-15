import { Amplify } from 'aws-amplify';
import { OpenAPI } from '$lib/generated/backend-client';
import { getIdentityToken } from '$lib/auth';

Amplify.configure({
    Auth: {
        Cognito: {
            //  Amazon Cognito User Pool ID
            userPoolId: 'eu-west-1_6LKveLbb7',
            // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
            userPoolClientId: '3o3qgc7btto12unlnbfio4m0no',
            // REQUIRED only for Federated Authentication - Amazon Cognito Identity Pool ID
            // identityPoolId: 'XX-XXXX-X:XXXXXXXX-XXXX-1234-abcd-1234567890ab',
            // OPTIONAL - This is used when autoSignIn is enabled for Auth.signUp
            // 'code' is used for Auth.confirmSignUp, 'link' is used for email link verification
            // signUpVerificationMethod: 'link', // 'code' | 'link'
            loginWith: {
                // OPTIONAL - Hosted UI configuration
                oauth: {
                    domain: 'auth.blog.butterneck.me',
                    scopes: [
                        'email',
                        'profile',
                        'openid',
                    ],
                    redirectSignIn: ['https://blog.butterneck.me'],
                    redirectSignOut: ['https://blog.butterneck.me'],
                    responseType: 'code' // or 'token', note that REFRESH token will only be generated when the responseType is code
                }
            }
        }
    }
});

OpenAPI.TOKEN = async () => getIdentityToken();
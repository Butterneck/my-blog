import { Amplify } from 'aws-amplify';
Amplify.configure({
    Auth: {
        Cognito: {
            //  Amazon Cognito User Pool ID
            userPoolId: 'eu-west-1_bWybVHw05',
            // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
            userPoolClientId: '1m1cvqhgc1a1eo9miqdlfah4l6',
            // REQUIRED only for Federated Authentication - Amazon Cognito Identity Pool ID
            // identityPoolId: 'XX-XXXX-X:XXXXXXXX-XXXX-1234-abcd-1234567890ab',
            // OPTIONAL - This is used when autoSignIn is enabled for Auth.signUp
            // 'code' is used for Auth.confirmSignUp, 'link' is used for email link verification
            signUpVerificationMethod: 'link', // 'code' | 'link'
            // loginWith: {
            //     // OPTIONAL - Hosted UI configuration
            //     oauth: {
            //         domain: 'your_cognito_domain',
            //         scopes: [
            //             'phone',
            //             'email',
            //             'profile',
            //             'openid',
            //             'aws.cognito.signin.user.admin'
            //         ],
            //         redirectSignIn: ['http://localhost:3000/'],
            //         redirectSignOut: ['http://localhost:3000/'],
            //         responseType: 'code' // or 'token', note that REFRESH token will only be generated when the responseType is code
            //     }
            // }
        }
    }
});
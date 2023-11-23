import {
    signIn as ampSignIn,
    type SignInInput,
    signOut as ampSignOut,
    getCurrentUser as ampGetCurrentUser,
    type AuthUser,
    fetchAuthSession as ampFetchAuthSession,
} from "aws-amplify/auth"

export async function logIn({ username, password }: SignInInput): Promise<boolean> {
    try {
        const { isSignedIn, nextStep } = await ampSignIn({ username, password });
        return isSignedIn;
    } catch (error) {
        console.error('error signing in', error);
    }

    return false;
}

export async function logOut(): Promise<void> {
    try {
        await ampSignOut();
    } catch (error) {
        console.error('error signing out', error);
    }
}

export async function getCurrentUser(): Promise<AuthUser | null> {
    try {
        const user = await ampGetCurrentUser();
        return user;
    } catch (error) {
        console.error('error getting current user', error);
    }

    return null;
}

export async function getIdentityToken(): Promise<string | null> {

    try {
        const { accessToken, idToken } = (await ampFetchAuthSession()).tokens ?? {};
        if (idToken) {
            return idToken.toString();
        }
    } catch (error) {
        console.error('error getting identity token', error);
    }

    return null;
}
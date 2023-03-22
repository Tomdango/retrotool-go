"use client";

import { Amplify, Auth } from "aws-amplify";
import {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from "react";

Amplify.configure({
  Auth: {
    region: process.env.COGNITO_AUTH_REGION,
    userPoolId: process.env.COGNITO_AUTH_USER_POOL_ID,
    userPoolWebClientId: process.env.COGNITO_AUTH_USER_CLIENT_ID,
    cookieStorage: {
      domain: process.env.COGNITO_AUTH_COOKIE_STORAGE_DOMAIN,
      path: "/",
      expires: 365,
      sameSite: "strict",
      secure: true,
    },
    authenticationFlowType: "USER_SRP_AUTH",
  },
});

type LoginResult = {
  success: boolean;
  message: string;
};

type IAuthContext = AuthState & {
  isLoading: boolean;
  login: (username: string, password: string) => Promise<LoginResult>;
  logout: () => Promise<LoginResult>;
};

const AuthContext = createContext<IAuthContext>({
  isLoggedIn: false,
  isLoading: false,
  login: async () => ({ success: false, message: "NOT_INITIALISED" }),
  logout: async () => ({ success: false, message: "NOT_INITIALISED" }),
});

export const useAuth = () => useContext(AuthContext);

interface LoggedOutAuthState {
  isLoggedIn: false;
}

interface LoggedInAuthState {
  isLoggedIn: true;
  details: {
    username: string;
    phoneNumber: string;
    phoneNumberVerified: boolean;
  };
  session: {
    idToken: string;
    accessToken: string;
    refreshToken: string;
  };
}

type AuthState = LoggedInAuthState | LoggedOutAuthState;

const useAuthProvider = (): IAuthContext => {
  const [isLoading, setIsLoading] = useState(true);
  const [authState, setAuthState] = useState<AuthState>({ isLoggedIn: false });

  useEffect(() => {
    Auth.currentAuthenticatedUser()
      .then((result) => {
        setIsLoading(false);
        console.log(result);
        setAuthState({
          isLoggedIn: true,
          details: {
            username: result.username,
            phoneNumber: result.attributes.phone_number,
            phoneNumberVerified: result.attributes.phone_number_verified,
          },
          session: {
            accessToken: result.signInUserSession.accessToken.jwtToken,
            refreshToken: result.signInUserSession.refreshToken.jwtToken,
            idToken: result.signInUserSession.idToken.jwtToken,
          },
        });
      })
      .catch(() => {
        setAuthState({ isLoggedIn: false });
        setIsLoading(false);
      });
  }, []);

  const login = async (username: string, password: string) => {
    try {
      const result = await Auth.signIn(username, password);
      setAuthState({
        isLoggedIn: true,
        details: {
          username: result.username,
          phoneNumber: result.attributes.phone_number,
          phoneNumberVerified: result.attributes.phone_number_verified,
        },
        session: {
          accessToken: result.signInUserSession.accessToken.jwtToken,
          refreshToken: result.signInUserSession.refreshToken.jwtToken,
          idToken: result.signInUserSession.idToken.jwtToken,
        },
      });
      console.log(result);
      return { success: true, message: "" };
    } catch (error) {
      console.error(error);
      return {
        success: false,
        message: "LOGIN FAIL",
      };
    }
  };

  const logout = async () => {
    try {
      await Auth.signOut();
      setAuthState({ isLoggedIn: false });
      return { success: true, message: "" };
    } catch (error) {
      return {
        success: false,
        message: "LOGOUT FAIL",
      };
    }
  };

  return {
    ...authState,
    isLoading,
    login,
    logout,
  };
};

export const AuthProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const auth = useAuthProvider();
  return <AuthContext.Provider value={auth}>{children}</AuthContext.Provider>;
};

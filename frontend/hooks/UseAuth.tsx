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

interface IAuthContext {
  isLoading: boolean;
  isAuthenticated: boolean;
  username: string;
  login: (username: string, password: string) => Promise<LoginResult>;
  logout: () => Promise<LoginResult>;
}

const AuthContext = createContext<IAuthContext>({
  isLoading: false,
  isAuthenticated: false,
  username: "NOT_INITIALISED",
  login: async () => ({ success: false, message: "NOT_INITIALISED" }),
  logout: async () => ({ success: false, message: "NOT_INITIALISED" }),
});

export const useAuth = () => useContext(AuthContext);

const useAuthProvider = (): IAuthContext => {
  const [isLoading, setIsLoading] = useState(true);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [username, setUsername] = useState("");

  useEffect(() => {
    Auth.currentAuthenticatedUser()
      .then((result) => {
        setUsername(result.username);
        setIsAuthenticated(true);
        setIsLoading(false);
      })
      .catch(() => {
        setUsername("");
        setIsAuthenticated(false);
        setIsLoading(false);
      });
  }, []);

  const login = async (username: string, password: string) => {
    try {
      const result = await Auth.signIn(username, password);
      setUsername(result.username);
      setIsAuthenticated(true);
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
      setUsername("");
      setIsAuthenticated(false);
      return { success: true, message: "" };
    } catch (error) {
      return {
        success: false,
        message: "LOGOUT FAIL",
      };
    }
  };

  return {
    isLoading,
    isAuthenticated,
    username,
    login,
    logout,
  };
};

export const AuthProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const auth = useAuthProvider();
  return <AuthContext.Provider value={auth}>{children}</AuthContext.Provider>;
};

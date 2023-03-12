"use client";

import { useAuth } from "@/hooks/UseAuth";
import axios from "axios";
import clsx from "clsx";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { FormEvent, useContext, useEffect, useState } from "react";

enum LoginStatus {
  NOT_STARTED,
  LOADING,
  FAILED,
  SUCCESS,
}

const LoginPage: React.FC = () => {
  const [status, setStatus] = useState(LoginStatus.NOT_STARTED);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const auth = useAuth();

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    setStatus(LoginStatus.LOADING);

    const result = await auth.login(username, password);

    if (result.success) {
      setStatus(LoginStatus.SUCCESS);
      router.prefetch("/");
      setTimeout(() => {
        router.replace("/");
      }, 3000);
    } else {
      setStatus(LoginStatus.FAILED);
    }
  };

  useEffect(() => {
    setStatus(LoginStatus.NOT_STARTED);
  }, [username, password]);

  return (
    <div className="hero fullscreen bg-green-600 ">
      <div className="hero-body">
        <div className="content">
          <h2 className="title text-white">👋 Welcome back!</h2>
          <h5 className="subtitle text-teal-300">
            Sign back in to get rolling.
          </h5>
          <div className="card p-3" style={{ maxWidth: 800 }}>
            <div
              className={clsx("toast toast--warning m-0", {
                "u-none": status !== LoginStatus.FAILED,
              })}
            >
              <p className="text-dark">
                <b>Login Failed.</b> Check your details and try again.
              </p>
            </div>
            <div
              className={clsx("toast toast--success m-0", {
                "u-none": status !== LoginStatus.SUCCESS,
              })}
            >
              <p>
                <b>Successfully logged in.</b> Welcome back {auth.username}!
              </p>
            </div>
            <form onSubmit={handleSubmit}>
              <div className="input-control">
                <input
                  value={username}
                  onChange={(e) => setUsername(e.currentTarget.value)}
                  type="username"
                  className="input-contains-icon"
                  placeholder="Username"
                />
                <span className="icon">
                  <i className="fa fa-wrapper fa-user"></i>
                </span>
              </div>
              <div className="input-control">
                <input
                  value={password}
                  onChange={(e) => setPassword(e.currentTarget.value)}
                  type="password"
                  className="input-contains-icon"
                  placeholder="Password"
                />
                <span className="icon">
                  <i className="fa fa-wrapper fa-lock"></i>
                </span>
              </div>
              <button
                disabled={status !== LoginStatus.NOT_STARTED}
                className={clsx({
                  "btn-success":
                    status === LoginStatus.NOT_STARTED ||
                    status === LoginStatus.SUCCESS,
                  "btn-dark": status === LoginStatus.LOADING,
                  "btn-warning": status === LoginStatus.FAILED,
                  "text-dark": status === LoginStatus.FAILED,
                })}
              >
                {status === LoginStatus.NOT_STARTED && "Log In"}
                {status === LoginStatus.LOADING && (
                  <>
                    <div className="u-flex u-items-center u-justify-center">
                      <div className="animated loading hide-text loading-white ml-1 mr-2"></div>
                      Logging in...
                    </div>
                  </>
                )}
                {status === LoginStatus.FAILED && "Login Failed"}
                {status === LoginStatus.SUCCESS && "Login Successful"}
              </button>
            </form>
            <div className="card__footer">
              <span>
                Need an account? <Link href="/auth/signup">Sign up here!</Link>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;

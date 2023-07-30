'use client';

import { login } from "@/graphql/auth/login";
import { ACCOUNT_HOME } from "@/lib/constants";
import { useTokenStore } from "@/lib/store";
import { gql, useMutation } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { ErrorParser } from "@/lib/errorParser";

type FormValues = {
    email: string;
    password: string;
};


export default function Login() {
    const { register, handleSubmit } = useForm<FormValues>();

    const navigate = useNavigate();

    const [executeLogin, { data, loading, error }] = useMutation(gql(login), { errorPolicy: "all" });

    const { setToken } = useTokenStore();

    const onSubmit: SubmitHandler<FormValues> = (formValues) => {
        executeLogin({
            variables: {
                email: formValues.email,
                password: formValues.password,
            },
        });
    };

    useEffect(() => {
        if (data?.authentication?.login?.jwt?.length > 0) {
            let newToken = data?.authentication?.login?.jwt;
            setToken(newToken);
            navigate(ACCOUNT_HOME);
        }
    }, [data, setToken]);

    return (
        <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                    Sign in to your account
                </h2>
            </div>

            <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                            Email address
                        </label>
                        <div className="mt-2">
                            <input
                                id="email"
                                type="email"
                                autoComplete="email"
                                required
                                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                {...register("email")}
                            />
                        </div>
                    </div>

                    <div>
                        <div className="flex items-center justify-between">
                            <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                                Password
                            </label>
                            <div className="text-sm">
                                <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">
                                    Forgot password?
                                </a>
                            </div>
                        </div>
                        <div className="mt-2">
                            <input
                                id="password"
                                type="password"
                                autoComplete="current-password"
                                required
                                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                {...register("password")}
                            />
                        </div>
                    </div>

                    <div>
                        <button
                            type="submit"
                            className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                        >
                            Sign in
                        </button>
                    </div>
                </form>
                {error && <p className="mt-10 text-center text-sm text-red-600 font-bold">
                    {ErrorParser(error.message)}
                </p>}
            </div>
        </div>
    )
}

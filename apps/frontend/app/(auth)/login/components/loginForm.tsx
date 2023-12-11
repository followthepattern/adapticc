import { gql, useMutation } from "@apollo/client";
import { login } from "@/graphql/auth/login";
import { SubmitHandler, useForm } from "react-hook-form";
import { useEffect } from "react";
import { useTokenStore } from "@/lib/store";
import { ErrorParser } from "@/lib/errorParser";

type FormValues = {
    email: string;
    password: string;
};

export default function LoginForm() {
    const { register, handleSubmit } = useForm<FormValues>();

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
            const newToken = data?.authentication?.login?.jwt;
            setToken(newToken);
        }
    }, [data, setToken]);

    const inputLabelClassName = "block text-sm font-medium text-gray-900"
    const inputFieldClassName = "block w-full py-2 border-0 rounded-md shadow-sm text-gray-950 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-blue-600 sm:text-sm";

    return (
        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                <div>
                    <label htmlFor="email" className={inputLabelClassName}>
                        Email address
                    </label>
                    <div className="mt-2">
                        <input
                            id="email"
                            type="email"
                            autoComplete="email"
                            required
                            className={inputFieldClassName}
                            {...register("email")}
                        />
                    </div>
                </div>

                <div>
                    <div className="flex items-center justify-between">
                        <label htmlFor="password" className={inputLabelClassName}>
                            Password
                        </label>
                    </div>
                    <div className="mt-2">
                        <input
                            id="password"
                            type="password"
                            autoComplete="current-password"
                            required
                            className={inputFieldClassName}
                            {...register("password")}
                        />
                    </div>
                </div>

                <div>
                    <button
                        type="submit"
                        className="w-full px-3 py-3 text-sm font-semibold text-white bg-blue-500 rounded-md shadow-sm mx:auto hover:bg-blue-700 focus:bg-blue-900"
                    >
                        Sign in
                    </button>
                    {error && <p className="mt-2 text-sm font-bold text-red-600">
                        {ErrorParser(error.message)}
                    </p>}
                </div>

                <div>

                </div>
            </form>
        </div>

    )
}
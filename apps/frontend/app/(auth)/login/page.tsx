import { ACCOUNT_HOME } from "@/lib/constants";
import { useTokenStore } from "@/lib/store";
import { Navigate } from "react-router-dom";
import LoginForm from "./components/loginForm";


export default function Login() {
    const { token } = useTokenStore();

    if (token != "") {
        return <Navigate to={ACCOUNT_HOME} replace={true} />
    }

    return (
        <div className="flex flex-col justify-center min-h-full px-6 py-12 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 className="mt-10 text-2xl font-bold text-center text-gray-900">
                    Sign in to your account!
                </h2>
            </div>
            <LoginForm />
        </div>
    )
}

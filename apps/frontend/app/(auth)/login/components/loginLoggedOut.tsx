import LoginForm from "./loginForm";

export default function LoginExpired() {
    return (
        <div className="flex flex-col justify-center min-h-full px-6 py-12 lg:px-8">
            <div className="sm:mx-auto sm:max-w-sm">
                <h2 className="mt-10 text-2xl font-bold text-center text-gray-950">
                    You are logged out. Please sign in!
                </h2>
            </div>
            <LoginForm />
        </div>
    )
}
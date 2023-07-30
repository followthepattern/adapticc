import WithGraphQL from "@/components/withGraphQL";
import { ACCOUNT_HOME } from "@/lib/constants";
import { useTokenStore } from "@/lib/store";
import { Outlet, useNavigate } from "react-router-dom";

interface AuthLayoutProperties {}

export default function AuthLayout(props: AuthLayoutProperties) {
    const { token } = useTokenStore();
    const navigate = useNavigate();

    if (token) {
        navigate(ACCOUNT_HOME);
    }

    return (
        <WithGraphQL>
            <Outlet />
        </WithGraphQL>
    )
}

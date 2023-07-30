import WithGraphQL from "@/components/withGraphQL";
import { ACCOUNT_HOME } from "@/lib/constants";
import { useTokenStore } from "@/lib/store";
import useHasMounted from "@/lib/useMounted";
import { Outlet, useNavigate } from "react-router-dom";

interface AuthLayoutProperties {}

export default function AuthLayout(props: AuthLayoutProperties) {
    const { token } = useTokenStore();
    const hasMounted = useHasMounted();
    const navigate = useNavigate();

    if (!hasMounted) {
        return null;
    }

    if (token) {
        navigate(ACCOUNT_HOME);
    }

    return (
        <WithGraphQL>
            <Outlet />
        </WithGraphQL>
    )
}

'use client';

import WithGraphQL from "@/components/withGraphQL";
import { ACCOUNT_HOME } from "@/lib/constants";
import { useTokenStore } from "@/lib/store";
import useHasMounted from "@/lib/useMounted";
import { redirect } from "next/navigation";

interface AuthLayoutProperties {
    children: React.ReactNode,
}

export default function AuthLayout({ children }: AuthLayoutProperties) {
    const { token } = useTokenStore();
    const hasMounted = useHasMounted();

    if (!hasMounted) {
        return null;
    }

    if (token) {
        redirect(ACCOUNT_HOME);
    }

    return (
        <WithGraphQL>
            {children}
        </WithGraphQL>
    )
}

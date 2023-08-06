import { getUserProfile } from "@/graphql/users/query";
import { gql, useLazyQuery } from "@apollo/client";
import { useEffect } from "react";
import UserContext from "./userContext";
import { useTokenStore } from "@/lib/store";
import { QueryResponse } from "@/graphql/query";

interface WithUserContextProperties {
    children?: any;
}

const WithUserContext = (props: WithUserContextProperties) => {
    const [executeGetUserProfile, { data, called, loading, error }] = useLazyQuery<QueryResponse>(gql(getUserProfile), { errorPolicy: "all"})

    const { removeToken } = useTokenStore();

    useEffect(() => {
        if (!called) {
            executeGetUserProfile();
        }
    }, [called, executeGetUserProfile]);

    if (error) {
        removeToken();
        return <div>Something went wrong...</div>
    }

    const profile = data?.users?.profile;

    if (!profile || loading) {
        return <div>Loading.</div>
    }

    return (
        <UserContext.Provider value={profile}>
            <div>{props.children}</div>
        </UserContext.Provider>

    )
}

export default WithUserContext;
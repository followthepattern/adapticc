import useSingle from "@/graphql/hooks/useSingle";
import { QueryResponse, SingleQueryResult } from "@/graphql/query";
import { getSingleUser as graphQL } from "@/graphql/users/query";
import { User } from "@/models/user";

export default function useSingleUser(id: string): SingleQueryResult<User> {
    const parseResult = (data?: QueryResponse): User | undefined => {
        return data?.users?.single
    }

    return useSingle<User>({id, parseResult, graphQL})
}
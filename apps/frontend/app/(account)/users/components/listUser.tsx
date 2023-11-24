import { getUsers as graphQL } from "@/graphql/users/query";
import { ListResponse, ListQueryResult } from "@/graphql/utils/list";
import { QueryResponse } from "@/graphql/query";
import useList from "@/graphql/hooks/useList";
import { User } from "@/models/user";

export default function useListUsers(): ListQueryResult<ListResponse<User>> {
    const parseResult = (data?: QueryResponse): ListResponse<User> | undefined => {
        return data?.users?.list;
    }

    return useList<User>({parseResult, graphQL})
}
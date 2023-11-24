import { QueryResponse, SingleQueryResult } from "@/graphql/query";
import { gql, useQuery } from "@apollo/client";

interface SingleProperties<T> {
    id: string
    graphQL: string
    parseResult: (data?: QueryResponse) => T | undefined
}

export default function useSingle<T>({id, graphQL, parseResult}: SingleProperties<T>): SingleQueryResult<T> {
    const { data, loading, error } = useQuery<QueryResponse>(gql(graphQL), {
        variables: {
            id: id,
        }
    });

    const single = parseResult(data)

    const itemNotFound = !loading && !single;

    return { data: single, loading, error, itemNotFound };
}
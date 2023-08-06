import { ListResponse, ListQueryParams, ListQueryResult } from "@/graphql/utils/list";
import { QueryResponse } from "@/graphql/query";
import { gql, useLazyQuery } from "@apollo/client";

interface ListProperties<T> {
    parseResult: (data?: QueryResponse) => ListResponse<T> | undefined
    graphQL: string
}

export default function useList<T>({parseResult, graphQL}: ListProperties<T>): ListQueryResult<ListResponse<T>> {
    const [executeQuery, { data, called, loading, error }] = useLazyQuery<QueryResponse>(gql(graphQL));

    const execute = ({ page, pageSize, search, orderBy }: ListQueryParams) => {
        executeQuery({
            variables: {
                page: page,
                pageSize: pageSize,
                search: search,
                orderBy: orderBy,
            }
        });
    };

    const list = parseResult(data);

    const itemNotFound = called && !loading && !list;

    return [execute, { data: list, called, loading, error, itemNotFound }];
}
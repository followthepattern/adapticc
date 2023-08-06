import { gql, useMutation } from "@apollo/client";
import { CreateMutationResult, MutationResponse } from "@/graphql/mutation";

interface CreateProperties {
    parseResult: (data?: MutationResponse | null) => number | undefined
    graphQL: string
}

export default function useCreate<T>({parseResult, graphQL}: CreateProperties): CreateMutationResult<T, number> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResponse>(gql(graphQL));

    const execute = (model: T) => {
        executeMutation({
            variables: {
                model: model
            }
        });
    };

    const result = parseResult(data);

    return [execute, { createResult: result, createLoading: loading, createError: error }];
}
import { gql, useMutation } from "@apollo/client";
import { MutationResponse, UpdateMutationResult } from "@/graphql/mutation";

interface UpdateProperties {
    parseResult: (data?: MutationResponse | null) => number | undefined
    graphQL: string
}

export default function useUpdate<T>({parseResult, graphQL}: UpdateProperties): UpdateMutationResult<T,number | undefined> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResponse>(gql(graphQL));

    const execute = (id: string, model: T) => {
        executeMutation({
            variables: {
                id: id,
                model: model
            }
        });
    };

    const code = parseResult(data);

    return [execute, { updateResult: code, updateLoading: loading, updateError: error }];
}
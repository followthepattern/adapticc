import { gql, useMutation } from "@apollo/client";
import { DeleteMutationResult, MutationResponse } from "@/graphql/mutation";

interface DeleteProperties {
    parseResult: (data?: MutationResponse | null) => number | undefined
    graphQL: string
}

export default function useDelete({graphQL, parseResult}: DeleteProperties): DeleteMutationResult<string,number | undefined> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResponse>(gql(graphQL));

    const execute = (id: string) => {
        executeMutation({
            variables: {
                id: id
            }
        });
    };

    const code = parseResult(data);

    return [execute, { deleteResult: code, deleteLoading: loading, deleteError: error }];
}
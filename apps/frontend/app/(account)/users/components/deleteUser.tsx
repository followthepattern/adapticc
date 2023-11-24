import { deleteUser as graphQL } from "@/graphql/users/mutation";
import { DeleteMutationResult, MutationResponse } from "@/graphql/mutation";
import useDelete from "@/graphql/hooks/useDelete";

export default function useDeleteUser(): DeleteMutationResult<string,number | undefined> {
    const parseResult = (data?: MutationResponse | null): number | undefined => {
        return data?.users?.delete?.code
    }

    return useDelete({graphQL, parseResult})
}
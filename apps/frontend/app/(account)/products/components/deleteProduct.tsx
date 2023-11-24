import { deleteProduct as graphQL } from "@/graphql/products/mutation";
import { DeleteMutationResult, MutationResponse } from "@/graphql/mutation";
import useDelete from "@/graphql/hooks/useDelete";

export default function useDeleteProduct(): DeleteMutationResult<string,number | undefined> {
    const parseResult = (data?: MutationResponse | null): number | undefined => {
        return data?.products?.delete?.code
    }

    return useDelete({graphQL, parseResult})
}
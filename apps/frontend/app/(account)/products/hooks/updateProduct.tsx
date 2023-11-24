import { updateProduct as graphQL } from "@/graphql/products/mutation";
import { MutationResponse, UpdateMutationResult } from "@/graphql/mutation";
import { Product } from "@/models/product";
import useUpdate from "@/graphql/hooks/useUpdate";

export default function useUpdateProduct(): UpdateMutationResult<Product,number | undefined> {
const parseResult = (data?: MutationResponse | null) : number | undefined  => {
        return data?.products?.update?.code;
    }

    return useUpdate({parseResult, graphQL})
}
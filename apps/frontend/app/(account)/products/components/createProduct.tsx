import { CreateMutationResult, MutationResponse } from "@/graphql/mutation";
import { Product } from "@/models/product";
import useCreate from "@/graphql/hooks/useCreate";
import { createProduct as graphQL } from "@/graphql/products/mutation";

export default function useCreateProduct(): CreateMutationResult<Product, number> {
    const parseResult = (data?: MutationResponse | null) : number | undefined  => {
        return data?.products?.create?.code;
    }

    return useCreate<Product>({parseResult, graphQL})
}
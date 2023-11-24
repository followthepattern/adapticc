import useSingle from "@/graphql/hooks/useSingle";
import { getSingleProduct as graphQL } from "@/graphql/products/query";
import { QueryResponse, SingleQueryResult } from "@/graphql/query";
import { Product } from "@/models/product";

export default function useSingleProduct(id: string): SingleQueryResult<Product> {
    const parseResult = (data?: QueryResponse): Product | undefined => {
        return data?.products?.single
    }

    return useSingle<Product>({id, parseResult, graphQL})
}
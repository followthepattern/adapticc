import { getProducts as graphQL } from "@/graphql/products/query";
import { ListResponse, ListQueryResult } from "@/graphql/utils/list";
import { QueryResponse } from "@/graphql/query";
import { Product } from "@/models/product";
import useList from "@/graphql/hooks/useList";

export default function useListProduct(): ListQueryResult<ListResponse<Product>> {
    const parseResult = (data?: QueryResponse): ListResponse<Product> | undefined => {
        return data?.products?.list;
    }

    return useList<Product>({parseResult, graphQL})
}
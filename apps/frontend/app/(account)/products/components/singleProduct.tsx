import { Product, getSingleProduct } from "@/graphql/products/query";
import { QueryResponse } from "@/graphql/query";
import { useQuery } from "@apollo/client";

interface SingleProductQueryResult<Data = any> {
      loading: boolean;
      data?: Data;
      error?: any;
      itemNotFound?: boolean;
}

export default function useSingleProduct(id: string): SingleProductQueryResult<Product> {
    const { data, loading, error } = useQuery<QueryResponse>(getSingleProduct, {variables: {
        id: id,
    }});

    const single = data?.products?.single;

    const itemNotFound = !loading && !single;

    return { data: single, loading, error, itemNotFound };
}
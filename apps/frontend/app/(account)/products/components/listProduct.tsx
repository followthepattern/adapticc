import { Product, getProducts } from "@/graphql/products/query";
import { List } from "@/graphql/utils/list";
import { QueryResponse } from "@/graphql/query";
import { useLazyQuery } from "@apollo/client";

type ListProductQueryResult<Data = any> = [
    (page: number, pageSize: number) => void,
    {
      loading: boolean;
      data?: Data;
      error?: any;
      itemNotFound?: boolean;
      called: boolean;
    }
  ];

export default function useListProduct(): ListProductQueryResult<List<Product>> {
    const [executeQuery, { data, called, loading, error }] = useLazyQuery<QueryResponse>(getProducts);

    const execute = (page: number, pageSize: number) => {
        executeQuery({
            variables: {
                page: page,
                pageSize: pageSize,
            }
        });
    };

    const list = data?.products?.list;

    const itemNotFound = called && !loading && !list;

    return [execute, { data: list, called, loading, error, itemNotFound }];
}
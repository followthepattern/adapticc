import { Product, getProducts } from "@/graphql/products/query";
import { List, ListFilter } from "@/graphql/utils/list";
import { QueryResponse } from "@/graphql/query";
import { useLazyQuery } from "@apollo/client";

type ListProductQueryResult<Data = any> = [
    (filter: ListFilter) => void,
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

    const execute = ({page, pageSize, search, orderBy}:ListFilter) => {
        executeQuery({
            variables: {
                page: page,
                pageSize: pageSize,
                search: search,
                orderBy: orderBy,
            }
        });
    };

    const list = data?.products?.list;

    const itemNotFound = called && !loading && !list;

    return [execute, { data: list, called, loading, error, itemNotFound }];
}
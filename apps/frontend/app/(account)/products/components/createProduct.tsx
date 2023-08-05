import { MutationResult, useMutation } from "@apollo/client";
import { ProductMutation, createProduct } from "@/graphql/products/mutation";
import { MutationResponse } from "@/graphql/mutation";

interface Product {
    id: string
    title: string
    description: string
}

type CreateProductMutationResult<Entity = any, TResult = any> = [
    (model: Entity) => void,
    {
        createLoading: boolean;
        createResult?: TResult;
        createError?: any;
    }
];

export default function useCreateProduct(): CreateProductMutationResult<Product, number | undefined> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResponse>(createProduct);

    const execute = (model: Product) => {
        executeMutation({
            variables: {
                model: model
            }
        });
    };

    
    const code = data?.products?.create?.code

    return [execute, { createResult: code, createLoading: loading, createError: error }];
}
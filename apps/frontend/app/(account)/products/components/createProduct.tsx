import { MutationResult, useMutation } from "@apollo/client";
import { ProductMutation, createProduct } from "@/graphql/products/mutation";

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
    const [executeMutation, { data, loading, error }] = useMutation<MutationResult<ProductMutation>>(createProduct);

    const execute = (model: Product) => {
        executeMutation({
            variables: {
                model: model
            }
        });
    };

    const code = data?.data?.update?.code

    return [execute, { createResult: code, createLoading: loading, createError: error }];
}
import { MutationResult, useMutation } from "@apollo/client";
import { ProductMutation, updateProduct } from "@/graphql/products/mutation";

interface Product {
    id: string
    title: string
    description: string
}

type UpdateProductMutationResult<Entity = any, TResult = any> = [
    (model: Entity) => void,
    {
        updateLoading: boolean;
        updateResult?: TResult;
        updateError?: any;
    }
];

export default function useUpdateProduct(): UpdateProductMutationResult<Product,number | undefined> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResult<ProductMutation>>(updateProduct);

    const execute = (model: Product) => {
        executeMutation({
            variables: {
                model: model
            }
        });
    };

    const code = data?.data?.update?.code

    return [execute, { updateResult: code, updateLoading: loading, updateError: error }];
}
import { MutationResult, useMutation } from "@apollo/client";
import { ProductMutation, deleteProduct, updateProduct } from "@/graphql/products/mutation";

type DeleteProductMutationResult<Entity = any, TResult = any> = [
    (id: Entity) => void,
    {
        deleteLoading: boolean;
        deleteResult?: TResult;
        deleteError?: any;
    }
];

export default function useDeleteProduct(): DeleteProductMutationResult<string,number | undefined> {
    const [executeMutation, { data, loading, error }] = useMutation<MutationResult<ProductMutation>>(deleteProduct);

    const execute = (id: string) => {
        executeMutation({
            variables: {
                id: id
            }
        });
    };

    const code = data?.data?.update?.code

    return [execute, { deleteResult: code, deleteLoading: loading, deleteError: error }];
}
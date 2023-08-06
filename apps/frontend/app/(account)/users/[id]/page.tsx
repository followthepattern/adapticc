import { MutationResponse } from "@/graphql/mutation";
import { QueryResponse } from "@/graphql/query";
import { deleteUser, updateUser } from "@/graphql/users/mutation";
import { getSingleUser } from "@/graphql/users/query";
import { gql, useLazyQuery, useMutation } from "@apollo/client";
import { useEffect } from "react";

interface ListPageWrapperProperties {
    params: {
        id: string
    }
}

function DeleteItem(id: string) {

}

interface UpdateItemButtonProperties {
    id: string
    firstName?: string
    lastName?: string
}

const UpdateItemButton = (props: UpdateItemButtonProperties) => {
    const [executeUpdateUser, { data, called, loading, error }] = useLazyQuery<MutationResponse>(gql(updateUser));

    // const isLoading = called && loading;

    return (
        <button onClick={() => executeUpdateUser()}>
            Update
        </button>
    )
}

export default function Page({ params: { id } }: ListPageWrapperProperties) {
    const [executeSingleUser, { data, called, loading, error }] = useLazyQuery<QueryResponse>(gql(getSingleUser))

    const [executeUpdateUser, { data: updateData, called: updateCalled, loading: updateLoading, error: updateError }] = useMutation<MutationResponse>(gql(updateUser));

    const [executeDeleteUser, { data: deleteData, called: deleteCalled, loading: deleteLoading, error: deleteError }] = useMutation<MutationResponse>(deleteUser, { errorPolicy: "all" });


    useEffect(() => {
        if (!called) {
            executeSingleUser({
                variables: {
                    id: id,
                }
            });
        }
    }, [called, executeSingleUser, id]);

    const user = data?.users?.single;

    if (loading) {
        return (
            <div>Loading...</div>
        )
    }

    if (called && !loading && !user) {
        // notFound();
    }

    return (
        <div>
            <div>{JSON.stringify(data)}</div>
            <button onClick={() => executeUpdateUser({
                variables: {
                    id: id,
                    firstName: "testFirstName4",
                    lastName: "testLastName4",
                }
            })}>
                Update
            </button>
            <button onClick={() => {
                executeDeleteUser({
                    variables: {
                        id: id,
                    }
                })
            }}>{deleteLoading ? "loading" : "Delete"}</button>
            {deleteError == null ? "" : <div>{JSON.stringify(deleteError)}</div>}
        </div>
    )
}
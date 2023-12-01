import useSingleProduct from "../../hooks/singleProduct";
import useUpdateProduct from "../../hooks/updateProduct";
import { useForm } from "react-hook-form";
import useDeleteProduct from "../../hooks/deleteProduct";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import { useEffect } from "react";
import { Product } from "@/models/product";
import SingleLayout from "@/app/(account)/components/singlePage/layout";
import CancelButton from "@/app/(account)/components/buttons/cancelButton";
import EditButton from "@/app/(account)/components/buttons/editButton";
import DeleteButton from "@/app/(account)/components/buttons/deleteButton";

export default function ProductEdit() {
    const { id } = useParams();

    const { loading, data, error, itemNotFound } = useSingleProduct(id ?? "");

    const [executeUpdateMutation, { updateLoading, updateError, updateResult }] = useUpdateProduct();

    const [executeDeleteMutation, { deleteLoading, deleteError, deleteResult }] = useDeleteProduct();

    const navigate = useNavigate();

    useEffect(() => {
        if (updateResult && updateResult > 0) {
            navigate(RESOURCE_URL);
        }
    }, [updateResult])

    useEffect(() => {
        if (deleteResult && deleteResult > 0) {
            navigate(RESOURCE_URL);
        }
    }, [deleteResult])

    const { register, getValues } = useForm<Product>();

    if (loading) {
        return (<div>Loading...</div>)
    }

    if (itemNotFound) {
        // notFound();
    }

    if (error) {
        return <div>{JSON.stringify(error)}</div>
    }

    const onSave = () => {
        const values = getValues();
        executeUpdateMutation(id ?? "", values)
    }

    const onDelete = () => {
        executeDeleteMutation(id ?? "")
    }

    const onCancel = () => {
        navigate(RESOURCE_URL)
    }

    return (
        <SingleLayout>
            <SingleLayout.Title>Product: {data?.id}</SingleLayout.Title>
            <div className="grid grid-cols-1 mt-10 gap-x-6 gap-y-8 sm:grid-cols-6">
                <div className="sm:col-span-3">
                    <label htmlFor="title" className="block text-sm font-medium leading-6 text-gray-900">
                        Title
                    </label>
                    <div className="mt-2">
                        <input
                            {...register("title")}
                            type="text"
                            id="title"
                            defaultValue={data?.title}
                            className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        />
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="description" className="block text-sm font-medium leading-6 text-gray-900">
                        Description
                    </label>
                    <div className="mt-2">
                        <textarea
                            id="description"
                            {...register("description")}
                            className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            defaultValue={data?.description}
                        />
                    </div>
                </div>
            </div>
            <SingleLayout.Footer className="justify-between">
                <DeleteButton onClick={onDelete} />
                <div className="flex gap-x-2">
                    <CancelButton onClick={onCancel} />
                    <EditButton onClick={onSave} />
                </div>
            </SingleLayout.Footer>
        </SingleLayout>
    )
}
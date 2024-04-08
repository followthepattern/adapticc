import useSingleProduct from "../../hooks/singleProduct";
import useUpdateProduct from "../../hooks/updateProduct";
import { useForm } from "react-hook-form";
import useDeleteProduct from "../../hooks/deleteProduct";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import React, { useEffect } from "react";
import { Product } from "@/models/product";
import SingleLayout from "@/app/(account)/components/singleView/layout";
import SecondaryButton from "@/app/(account)/components/buttons/secondaryButton";
import PrimaryButton from "@/app/(account)/components/buttons/primaryButton";
import GridFields from "@/app/(account)/components/singleView/gridFields/gridFields";
import Label from "@/app/(account)/components/labels/label";
import Input from "@/app/(account)/components/inputFields/input";
import TextArea from "@/app/(account)/components/inputFields/textarea";
import { Id, toast } from "react-toastify";
import ConfirmModal from "@/app/(account)/components/modals/confirmModal";

export default function ProductEdit() {
    const { id } = useParams();

    const { loading, data, error, itemNotFound } = useSingleProduct(id ?? "");

    const [executeUpdateMutation, { updateError, updateResult }] = useUpdateProduct();

    const [executeDeleteMutation, { deleteError, deleteResult }] = useDeleteProduct();

    const navigate = useNavigate();

    const toastId = React.useRef<Id | null>(null);

    useEffect(() => {
        if (updateResult && updateResult > 0) {
            if (toastId.current) {
                toast.update(toastId.current, {
                    render: "Success!",
                    type: toast.TYPE.SUCCESS,
                    autoClose: 3000,
                })
            }
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

        toastId.current = toast("Saving...", { autoClose: false })

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
            <form action="">
                <GridFields className="py-6">
                    <div className="sm:col-span-2">
                        <Label htmlFor="title">
                            Title
                        </Label>
                        <Input
                            {...register("title")}
                            type="text"
                            id="title"
                            defaultValue={data?.title}
                        />
                    </div>
                    <div className="sm:col-span-2">
                        <Label htmlFor="description">
                            Description
                        </Label>
                        <TextArea
                            id="description"
                            {...register("description")}
                            defaultValue={data?.description}
                        />
                    </div>
                </GridFields>
                <SingleLayout.Footer className="justify-between">
                    <div className="flex gap-x-2">
                        <PrimaryButton onClick={onSave}>Save</PrimaryButton>
                        <SecondaryButton onClick={onCancel}>Cancel</SecondaryButton>
                    </div>
                    <ConfirmModal onConfirm={onDelete} title="Delete products" body={`Are you sure you want to delete ${data?.title}?`}>
                        Delete
                    </ConfirmModal>
                </SingleLayout.Footer>
            </form>
        </SingleLayout>
    )
}
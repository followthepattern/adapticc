import useSingleProduct from "../../hooks/singleProduct";
import useUpdateProduct from "../../hooks/updateProduct";
import { useForm } from "react-hook-form";
import useDeleteProduct from "../../hooks/deleteProduct";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import { useEffect } from "react";
import { Product } from "@/models/product";
import SingleLayout from "@/app/(account)/components/singleView/layout";
import SecondaryButton from "@/app/(account)/components/buttons/secondaryButton";
import PrimaryButton from "@/app/(account)/components/buttons/primaryButton";
import AlertButton from "@/app/(account)/components/buttons/alertButton";
import GridFields from "@/app/(account)/components/gridFields/gridFields";
import Label from "@/app/(account)/components/labels/label";
import classNames from "classnames";
import Input from "@/app/(account)/components/inputs/input";
import TextArea from "@/app/(account)/components/inputs/textarea";

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
            <form>
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
                    <AlertButton onClick={onDelete}>
                        Delete
                    </AlertButton>
                    <div className="flex gap-x-2">
                        <SecondaryButton onClick={onCancel}>Cancel</SecondaryButton>
                        <PrimaryButton onClick={onSave}>Save</PrimaryButton>
                    </div>
                </SingleLayout.Footer>
            </form>
        </SingleLayout>
    )
}
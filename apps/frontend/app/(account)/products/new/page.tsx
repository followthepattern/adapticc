import { useForm } from "react-hook-form";
import useCreateProduct from "../hooks/createProduct";
import { useNavigate } from "react-router-dom";
import { RESOURCE_URL } from "../page";
import { useEffect } from "react";
import { Product } from "@/models/product";
import SingleLayout from "../../components/singleView/layout";
import GridFields from "../../components/singleView/gridFields/gridFields";
import Label from "../../components/labels/label";
import Input from "../../components/inputFields/input";
import TextArea from "../../components/inputFields/textarea";
import SecondaryButton from "../../components/buttons/secondaryButton";
import PrimaryButton from "../../components/buttons/primaryButton";

export default function ProductNew() {
    const [executeMutation, { createLoading, createError, createResult }] = useCreateProduct();

    const navigate = useNavigate();

    useEffect(() => {
        if (createResult && createResult > 0) {
            navigate(RESOURCE_URL);
        }

    }, [createResult])

    const { register, getValues } = useForm<Product>();

    const onCreate = () => {
        const values = getValues();

        executeMutation({
            id: "",
            title: values.title,
            description: values.description,
        })
    }

    const onCancel = () => {
        navigate(RESOURCE_URL)
    }

    return (
        <SingleLayout>
            <SingleLayout.Title>New Product</SingleLayout.Title>
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
                        />
                    </div>
                    <div className="sm:col-span-2">
                        <Label htmlFor="description">
                            Description
                        </Label>
                        <TextArea
                            id="description"
                            {...register("description")}
                        />
                    </div>
                </GridFields>
                <SingleLayout.Footer className="justify-end">
                    <SecondaryButton onClick={onCancel}>Cancel</SecondaryButton>
                    <PrimaryButton onClick={onCreate}>Create</PrimaryButton>
                </SingleLayout.Footer>
            </form>
        </SingleLayout>
    )
}
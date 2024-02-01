import { useForm } from "react-hook-form";
import useCreateProduct from "../hooks/createUser";
import { useNavigate } from "react-router-dom";
import { RESOURCE_URL } from "../page";
import { useEffect } from "react";
import { User } from "@/models/user";
import SingleLayout from "../../components/singleView/layout";
import GridFields from "../../components/singleView/gridFields/gridFields";
import Label from "../../components/labels/label";
import Input from "../../components/inputFields/input";
import SecondaryButton from "../../components/buttons/secondaryButton";
import PrimaryButton from "../../components/buttons/primaryButton";

export default function UserNew() {
    const [executeMutation, { createLoading, createError, createResult }] = useCreateProduct();

    const navigate = useNavigate();

    useEffect(() => {
        if (createResult && createResult > 0) {
            navigate(RESOURCE_URL);
        }

    }, [createResult])

    const { register, getValues } = useForm<User>();

    const onCreate = () => {
        const values = getValues();

        console.info(values)

        executeMutation(values);
    }

    const onCancel = () => {
        navigate(RESOURCE_URL)
    }

    return (
        <SingleLayout>
            <SingleLayout.Title>New User</SingleLayout.Title>
            <form>
                <GridFields className="py-6">
                    <div className="sm:col-span-2">
                        <Label htmlFor="title">
                            Email
                        </Label>
                        <Input
                            {...register("email")}
                            type="text"
                            id="email"
                        />
                    </div>
                    <div className="sm:col-span-2">
                        <Label htmlFor="firstName">
                            First Name
                        </Label>
                        <Input
                            id="firstName"
                            type="text"
                            {...register("firstName")}
                        />
                    </div>
                    <div className="sm:col-span-2">
                        <Label htmlFor="lastName">
                            Last Name
                        </Label>
                        <Input
                            id="lastName"
                            type="text"
                            {...register("lastName")}
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
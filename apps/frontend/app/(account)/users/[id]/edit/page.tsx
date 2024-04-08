import useSingleUser from "../../hooks/singleUser";
import useUpdateUser from "../../hooks/updateUser";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import React, { useEffect } from "react";
import useDeleteUser from "../../hooks/deleteUser";
import { User } from "@/models/user";
import SingleLayout from "@/app/(account)/components/singleView/layout";
import GridFields from "@/app/(account)/components/singleView/gridFields/gridFields";
import Label from "@/app/(account)/components/labels/label";
import Input from "@/app/(account)/components/inputFields/input";
import SecondaryButton from "@/app/(account)/components/buttons/secondaryButton";
import PrimaryButton from "@/app/(account)/components/buttons/primaryButton";
import { Id, toast } from 'react-toastify';
import ConfirmModal from "@/app/(account)/components/modals/confirmModal";

export default function UserEdit() {
    const { id } = useParams();

    const { loading, data, error, itemNotFound } = useSingleUser(id ?? "");

    const [executeUpdateMutation, { updateError, updateResult }] = useUpdateUser();

    const [executeDeleteMutation, { deleteError, deleteResult }] = useDeleteUser();

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

    const { register, getValues } = useForm<User>();

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

        toastId.current = toast("Saving...", { autoClose: false });

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
            <SingleLayout.Title>User: {data?.id}</SingleLayout.Title>
            <form action="">
                <GridFields className="py-6">
                    <div className="sm:col-span-2">
                        <Label htmlFor="title">
                            Email
                        </Label>
                        <Input
                            {...register("email")}
                            disabled={true}
                            type="text"
                            id="email"
                            defaultValue={data?.email}
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
                            defaultValue={data?.firstName}
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
                            defaultValue={data?.lastName}
                        />
                    </div>
                </GridFields>
                <SingleLayout.Footer className="justify-between">
                    <div className="flex gap-x-2">
                        <PrimaryButton onClick={onSave}>Save</PrimaryButton>
                        <SecondaryButton onClick={onCancel}>Cancel</SecondaryButton>
                    </div>
                    <ConfirmModal onConfirm={onDelete} title="Delete users" body={`Are you sure you want to delete ${data?.firstName} ${data?.lastName}?`}>
                        Delete
                    </ConfirmModal>
                </SingleLayout.Footer>
            </form>

        </SingleLayout>
    )
}
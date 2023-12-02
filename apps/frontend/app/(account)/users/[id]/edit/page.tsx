import useSingleUser from "../../hooks/singleUser";
import useUpdateUser from "../../hooks/updateUser";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import { useEffect } from "react";
import useDeleteUser from "../../hooks/deleteUser";
import { User } from "@/models/user";
import SingleLayout from "@/app/(account)/components/singlePage/layout";
import GridFields from "@/app/(account)/components/gridFields/gridFields";
import Label from "@/app/(account)/components/labels/label";
import Input from "@/app/(account)/components/inputs/input";
import TextArea from "@/app/(account)/components/inputs/textarea";
import AlertButton from "@/app/(account)/components/buttons/alertButton";
import SecondaryButton from "@/app/(account)/components/buttons/secondaryButton";
import PrimaryButton from "@/app/(account)/components/buttons/primaryButton";

export default function UserEdit() {
    const { id } = useParams();

    const { loading, data, error, itemNotFound } = useSingleUser(id ?? "");

    const [executeUpdateMutation, { updateLoading, updateError, updateResult }] = useUpdateUser();

    const [executeDeleteMutation, { deleteLoading, deleteError, deleteResult }] = useDeleteUser();

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
            <form>
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
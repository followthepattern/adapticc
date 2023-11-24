import useSingleUser from "../../hooks/singleUser";
import useUpdateUser from "../../hooks/updateUser";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../../page";
import { useEffect } from "react";
import useDeleteUser from "../../hooks/deleteUser";
import { User } from "@/models/user";

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
        <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
            <form action="">
                <div className="space-y-12">
                    <div className="border-b border-gray-900/10 pb-12">
                        <h2 className="text-base font-semibold leading-7 text-gray-900">{data?.id}</h2>

                        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                            <div className="sm:col-span-3">
                                <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                                    Email
                                </label>
                                <div className="mt-2">
                                    <input
                                        type="text"
                                        id="email"
                                        defaultValue={data?.email}
                                        disabled
                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 disabled:cursor-not-allowed disabled:bg-gray-50 disabled:text-gray-500 disabled:ring-gray-200"
                                    />
                                </div>
                            </div>

                            <div className="col-span-full">
                                <label htmlFor="first-name" className="block text-sm font-medium leading-6 text-gray-900">
                                    First Name
                                </label>
                                <div className="mt-2">
                                    <textarea
                                        id="first-name"
                                        {...register("firstName")}
                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                        defaultValue={data?.firstName}
                                    />
                                </div>
                            </div>
                            <div className="col-span-full">
                                <label htmlFor="last-name" className="block text-sm font-medium leading-6 text-gray-900">
                                    Last Name
                                </label>
                                <div className="mt-2">
                                    <textarea
                                        id="last-name"
                                        {...register("lastName")}
                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                        defaultValue={data?.lastName}
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div className="mt-6 flex items-center justify-between gap-x-6">
                    <button
                        type="button"
                        onClick={onDelete}
                        className="rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-600"
                    >
                        Delete
                    </button>
                    <div className="flex gap-x-6">
                        <button
                            type="button"
                            className="text-sm font-semibold leading-6 text-gray-900"
                            onClick={onCancel}
                        >
                            Cancel
                        </button>
                        <button
                            type="button"
                            onClick={onSave}
                            className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                        >
                            Save
                        </button>
                    </div>
                </div>
            </form>
        </div>
    )
}
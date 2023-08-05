'use client';

import { useForm } from "react-hook-form";
import useCreateProduct from "../components/createProduct";
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../page";
import { useEffect } from "react";

interface PageProperties { }

type ProductValues = {
    title: string;
    description: string;
}


export default function ProductNew({ }: PageProperties) {
    const [executeMutation, { createLoading, createError, createResult }] = useCreateProduct();

    const navigate = useNavigate();

    useEffect(() => {
        if (createResult && createResult > 0) {
            navigate(RESOURCE_URL);
        }

    }, [createResult])

    const { register, getValues } = useForm<ProductValues>();

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
        <form action="">
            <div className="space-y-12">
                <div className="border-b border-gray-900/10 pb-12">
                    <h2 className="text-base font-semibold leading-7 text-gray-900">New Product</h2>

                    <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                        <div className="sm:col-span-3">
                            <label htmlFor="title" className="block text-sm font-medium leading-6 text-gray-900">
                                Title
                            </label>
                            <div className="mt-2">
                                <input
                                    {...register("title")}
                                    type="text"
                                    id="title"
                                    defaultValue=""
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
                                    defaultValue=""
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div className="mt-6 flex items-center justify-end gap-x-6">
                <button type="button" className="text-sm font-semibold leading-6 text-gray-900" onClick={onCancel}>
                    Cancel
                </button>
                <button
                    type="button"
                    onClick={onCreate}
                    className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                    Create
                </button>
            </div>
        </form>
    )
}
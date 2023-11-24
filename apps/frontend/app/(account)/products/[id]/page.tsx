import { useNavigate, useParams } from "react-router-dom";
import useSingleProduct from "../hooks/singleProduct";
import { RESOURCE_URL } from "../page";

export default function Product() {
    const { id } = useParams()
    const navigate = useNavigate();

    const { loading, data, error, itemNotFound } = useSingleProduct(id ?? "");

    if (loading) {
        return (
            <div>Loading...</div>
        )
    }

    if (error) {
        return (
            <div>{JSON.stringify(error)}</div>
        )
    }

    if (itemNotFound) {
        // notFound();
    }

    const onCancel = () => {
        navigate(RESOURCE_URL);
    }

    const onEdit = () => {
        navigate(`${RESOURCE_URL}/${id}/edit`);
    }

    return (
        <div>
            <div className="px-4 sm:px-0">
                <h3 className="text-base font-semibold leading-7 text-gray-900">Product</h3>
            </div>
            <div className="mt-6 border-y border-gray-100">
                <dl className="divide-y divide-gray-100">
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-gray-900">ID</dt>
                        <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{data?.id}</dd>
                    </div>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-gray-900">Title</dt>
                        <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{data?.title}</dd>
                    </div>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-gray-900">Description</dt>
                        <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{data?.description}</dd>
                    </div>
                </dl>
            </div>
            <div className="mt-6 flex items-center justify-end gap-x-6">
                <button type="button" className="text-sm font-semibold leading-6 text-gray-900" onClick={onCancel}>
                    Cancel
                </button>
                <button
                    type="button"
                    onClick={onEdit}
                    className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                    Edit
                </button>
            </div>
        </div>
    )
}
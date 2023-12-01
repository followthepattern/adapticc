export default function EditButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
    return (
        <button
            type="button"
            className="px-4 py-2 font-semibold text-blue-500 border border-blue-500 rounded-lg hover:bg-gray-100 focus:bg-gray-200"
            {...PushSubscriptionOptions}
        >
            Edit
        </button>
    )
}
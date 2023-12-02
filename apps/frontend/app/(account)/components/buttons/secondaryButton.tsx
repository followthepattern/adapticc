export default function SecondaryButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
    return (
        <button
            type="button"
            className="px-3 py-2 font-semibold text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-100 focus:bg-gray-200"
            {...props}
        >
            {props.children}
        </button>
    )
}
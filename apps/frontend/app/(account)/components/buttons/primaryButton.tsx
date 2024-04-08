export default function PrimaryButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
    return (
        <button
            type="button"
            className="px-4 py-2 font-semibold text-white bg-blue-500 border border-blue-500 rounded-lg hover:bg-blue-600 focus:bg-blue-700"
            {...props}
        >
            {props.children}
        </button>
    )
}
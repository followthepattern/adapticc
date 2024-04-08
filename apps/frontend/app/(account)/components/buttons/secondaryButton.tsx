export default function SecondaryButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
    return (
        <button
            type="button"
            className="px-3 py-2 font-semibold text-blue-500 border border-blue-500 rounded-lg hover:bg-blue-50 focus:bg-blue-100"
            {...props}
        >
            {props.children}
        </button>
    )
}
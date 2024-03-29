import classNames from "classnames"

const className = "px-4 py-2 font-semibold text-red-500 border border-red-500 rounded-lg hover:bg-red-100 focus:bg-red-200"

export default function AlertButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
    return (
        <button
            type="button"
            {...props}
            className={classNames(props.className, className)}
        >
            {props.children}
        </button>
    )
}

AlertButton.ClassName = className
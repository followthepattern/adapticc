import classNames from "classnames"

const className = "px-4 py-2 text-red-500 underline rounded-lg underline-offset-2 hover:bg-red-50 focus:bg-red-100"

export default function TertiaryButton(props: React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
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

TertiaryButton.ClassName = className
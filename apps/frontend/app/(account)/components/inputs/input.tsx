import classNames from "classnames";

export default function Input(props: React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement>) {
    return (
        <input
            {...props}
            className={classNames(
                props.className,
                "block w-full mt-2 border border-gray-300 rounded-lg ring-1 ring-inset ring-gray-100 placeholder:text-gray-400 focus:ring-0 focus:ring-inset",
                {
                    "cursor-not-allowed bg-gray-100 text-gray-600": props.disabled
                }
            )}
        />
    )
}
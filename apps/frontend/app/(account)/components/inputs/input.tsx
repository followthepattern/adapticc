import classNames from "classnames";

export default function Input(props: React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement>) {
    return (
        <input
            {...props}
            className={classNames(props.className, "block w-full py-2 mt-2 text-gray-900 border border-gray-300 rounded-lg ring-1 ring-inset ring-gray-100 placeholder:text-gray-400 focus:ring-0 focus:ring-inset")}
        />
    )
}
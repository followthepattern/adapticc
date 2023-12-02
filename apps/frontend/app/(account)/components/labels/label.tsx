import classNames from "classnames"

export default function Label(props: React.DetailedHTMLProps<React.LabelHTMLAttributes<HTMLLabelElement>, HTMLLabelElement>) {
    return (
        <label
            className={classNames(props.className, "block text-sm font-medium text-gray-900")}
        >
            {props.children}
        </label>
    )
}
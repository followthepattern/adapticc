import classNames from "classnames"

interface GridFieldsProperties {
    children?: any
    className?: string
}

export default function GridFields(props: GridFieldsProperties) {
    return (
        <div className={classNames(props.className, "grid grid-cols-1 gap-y-8 sm:grid-cols-3")}>
            {props.children}
        </div>
    )
} 
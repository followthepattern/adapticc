import classNames from "classnames"

interface DataRowProperties {
    name: any
    children: any
    classNames?: string

}

function DataRow(props: DataRowProperties) {
    return (
        <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <div className="text-sm font-medium">{props.name}</div>
            <div className="mt-1 text-sm text-gray-700 sm:col-span-2 sm:mt-0">{props.children}</div>
        </div>
    )
}

interface DataListProperties {
    children: any
    className?: string
}

export default function DataList(props: DataListProperties) {
    return (
        <div className={classNames(props.className, "border-gray-100 border-y")}>
            <div className="divide-y divide-gray-100">
                {props.children}
            </div>
        </div>
    )
}

DataList.Row = DataRow
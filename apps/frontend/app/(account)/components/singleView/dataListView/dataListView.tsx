import classNames from "classnames"

function DataRow(props: React.HTMLAttributes<HTMLDivElement>) {
    return (
        <div className={classNames(props.className, "px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0")}>
            {props.children}
        </div>
    )
}

function Label(props: React.HTMLAttributes<HTMLDivElement>) {
    return (
        <div {...props} className={classNames(props.className, "text-sm font-medium")}>{props.children}</div>
    )
}

function Field(props: React.HTMLAttributes<HTMLDivElement>) {
    return (
        <div className={classNames(props.className, "mt-1 text-sm text-gray-700 sm:col-span-2 sm:mt-0")}>{props.children}</div>
    )
}

interface DataListViewProperties {
    children: any
    className?: string
}

export default function DataListView(props: DataListViewProperties) {
    return (
        <div className={classNames(props.className, "border-gray-100 border-y")}>
            <div className="divide-y divide-gray-100">
                {props.children}
            </div>
        </div>
    )
}

DataListView.Label = Label
DataListView.Field = Field
DataListView.Row = DataRow
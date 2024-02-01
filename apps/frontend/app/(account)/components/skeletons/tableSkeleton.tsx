import classNames from "classnames"

export default function TableSkeleton() {
    return (
        <div className="grid p-4 border border-gray-200 divide-y divide-gray-200 rounded gap-y-4 animate-pulse">
            <SkeletonRow />
            <SkeletonRow className="pt-4" />
            <SkeletonRow className="pt-4" />
            <SkeletonRow className="pt-4" />
            <SkeletonRow className="pt-4" />
        </div>
    )
}

function SkeletonRow(props: React.HtmlHTMLAttributes<HTMLDivElement>) {
    return (
        <div className={classNames(props.className,"flex items-center justify-between")}>
            <div>
                <div className="w-40 h-2.5 mb-2 bg-gray-300 rounded-full"></div>
                <div className="w-32 h-2 bg-gray-200 rounded-full "></div>
            </div>
            <div className="h-2.5 bg-gray-300 rounded-full w-12"></div>
        </div>
    )
}
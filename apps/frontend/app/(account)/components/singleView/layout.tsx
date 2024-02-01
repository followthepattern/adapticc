import classNames from "classnames"

interface SingleLayoutTitleProperties {
    className?: string
    children?: any
}

function Title(props: SingleLayoutTitleProperties) {
    return <h3 className={classNames("font-semibold", props.className)}>{props.children}</h3>
}

function Subtitle(props: SingleLayoutTitleProperties) {
    return <p className={classNames("text-sm text-gray-500", props.className)}>{props.children}</p>
}

interface SingleLayoutFooterProperties {
    children?: any
    className?: string
}

function Footer(props: SingleLayoutFooterProperties) {
    return (
        <div className={classNames(props.className, "flex mt-6 gap-x-2")}>
            {props.children}
        </div>)
}

interface SingleLayoutProperties {
    children?: any
}

export default function SingleLayout(props: SingleLayoutProperties) {
    return (
        <div className="overflow-hidden">
            <div className="max-w-4xl mx-auto sm:px-6 lg:px-8">
                <div className="p-4 border border-gray-300 rounded-lg">
                    {props.children}
                </div>
            </div>
        </div>
    )
}

SingleLayout.Title = Title
SingleLayout.Subtitle = Subtitle
SingleLayout.Footer = Footer
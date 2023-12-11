import PlusIcon from "@/app/icons/plus";
import classNames from "classnames";
import { Link } from "react-router-dom";

interface NewResourceLinkProperties {
    resourceUrl: string;
    className?: string;
}

export default function NewResourceLink(props: NewResourceLinkProperties) {
    return (
        <Link
            className={classNames(props.className, "flex justify-center items-center gap-x-2 font-semibold bg-blue-500 text-white border-0 rounded-lg hover:bg-blue-700 focus:bg-blue-900")}
            to={`${props.resourceUrl}/new`}
        >
            <PlusIcon />
            <span>New</span>
        </Link>
    )
}
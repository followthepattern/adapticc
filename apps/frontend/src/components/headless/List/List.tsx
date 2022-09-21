import Pagination from "../Pagination/Pagination";
import Filters from "./Filters";
import Table from "../Table/Table";
import ActionsDropdown from "./ActionDropdown";
import ActionButton from "./ActionButton";

export interface ListProperties {
    children: any;
    className?: string;
}

const List = ({children, className}: ListProperties) => {
    return <div className={className}>{children}</div>
}

List.Filters = Filters;
List.Table = Table;
List.Pagination = Pagination;
List.ActionDropdown = ActionsDropdown;
List.ActionButton = ActionButton;

export default List;
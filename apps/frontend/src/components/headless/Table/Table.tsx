import Body from "./Body";
import Header from "./Header";

interface TableProperties {
    children?: any;
    className?: string;
}

const Table = ({children, className}: TableProperties) => {
    return (
        <table className={className}>
            {children}
        </table>
    )
}

Table.Header = Header;
Table.Body = Body;

export default Table;
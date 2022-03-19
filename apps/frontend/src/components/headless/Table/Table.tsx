interface TableProperties {
    children?: any;
}

const Table = ({children}: TableProperties) => {
    return (
        <td>
            {children}
        </td>
    )
}

export default Table;
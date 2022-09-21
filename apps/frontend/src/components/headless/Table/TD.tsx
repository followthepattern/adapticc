interface TDProperties {
    children?: any;
    className?: string;
}

const TD = ({children, className}: TDProperties) => {
    return (
        <td className={className}>
            {children}
        </td>
    )
}

export default TD;
interface THProperties {
    children?: any;
    className?: string;
}

const TH = ({children, className}: THProperties) => {
    return (
        <th className={className}>
            {children}
        </th>
    )
}

export default TH;
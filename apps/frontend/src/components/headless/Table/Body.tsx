import TR from "./TR";

interface BodyProperties {
    children?: any;
    className?: string;
}

const Body = ({children, className}: BodyProperties) => {
    return (
        <tbody className={className}>
            {children}
        </tbody>
    )
}

Body.TR = TR;

export default Body;
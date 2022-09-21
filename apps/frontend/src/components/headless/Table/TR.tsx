import TD from "./TD";
import TH from "./TH";

interface TRProperties {
    children?: any;
    className?: string;
    trKey?: any;
}

const TR = ({children, className, trKey}: TRProperties) => {
    return (
        <tr key={trKey} className={className}>
            {children}
        </tr>
    )
}

TR.TH = TH;
TR.TD = TD;

export default TR;
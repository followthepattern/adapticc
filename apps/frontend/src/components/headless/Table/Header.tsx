import TR from "./TR";

interface HeaderProperties {
    className?: string;
    children?: any;
}

const Header = ({
    children,
    className
}: HeaderProperties) => {
    return <thead className={className}>{children}</thead>
}

Header.TR = TR;

export default Header;
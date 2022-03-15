import { useNavigate } from "react-router-dom";

interface NavbarAccountProfileProperties {
  children?: any;
  className?: string;
  accountProfilePath: string;
  icon: JSX.Element
}

export const AccountProfile = ({
  children,
  className,
  accountProfilePath,
  icon
}: NavbarAccountProfileProperties) => {
  const navigate = useNavigate()
  return (
    <div className={className}>
      {icon}
      <div onClick={() => navigate(accountProfilePath)}>
        {children}
      </div>
    </div>
  );
};
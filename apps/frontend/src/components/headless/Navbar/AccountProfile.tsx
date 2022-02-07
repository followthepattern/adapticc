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
  return (
    <div className={className}>
      {icon}
      <a href={accountProfilePath}>
        {children}
      </a>
    </div>
  );
};
interface NavbarHeaderProperties {
  children?: any;
  className?: string;
}

const NavbarHeader = ({
  children,
  className,
}: NavbarHeaderProperties) => {
  return (
    <div className={className}>
        {children}
    </div>
  );
};

export default NavbarHeader;

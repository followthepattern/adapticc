const NavbarTitle = ({
  children,
  className,
}: {
  children: any;
  className: string;
}) => {
  return (
      <div className={className}>{children}</div>
  );
};

export default NavbarTitle;

export interface NavbarToggleProperties {
  children?: string,
  className?: string,
  onClick?: () => void,
}

const NavbarToggle = ({
  children,
  className,
  onClick,
}: NavbarToggleProperties) => {
  return (
    <div className={className} onClick={onClick}>{children}</div>
  );
};

export default NavbarToggle;

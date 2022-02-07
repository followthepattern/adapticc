export interface NavbarToggleProperties {
  className?: string,
  onClick?: () => void,
  icon: JSX.Element
}

const NavbarToggle = ({
  className,
  onClick,
  icon
}: NavbarToggleProperties) => {
  return (
    <div className={className} onClick={onClick}>{icon}</div>
  );
};

export default NavbarToggle;

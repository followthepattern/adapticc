interface NavbarProperties {
  children: any;
  className: string;
}

export const Navbar = ({children, className}: NavbarProperties) => {
    return (
      <div className={className}>{children}</div>
    );
  };
  
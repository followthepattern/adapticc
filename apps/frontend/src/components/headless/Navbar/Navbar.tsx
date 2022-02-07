import { Header } from "./NavbarHeader";
import { AccountProfile } from "./AccountProfile"

interface NavbarProperties {
  children: any;
  className: string;
}

const Navbar = ({ children, className }: NavbarProperties) => {
  return <div className={className}>{children}</div>;
};

Navbar.Header = Header;
Navbar.AccountProfile = AccountProfile;

export default Navbar;

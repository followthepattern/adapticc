import NavbarToggle from "../../../components/headless/Navbar/NavbarToggle";
import { useNavbarStore } from "../../../utils/store";
import { MenuIcon } from "@heroicons/react/outline";

interface HeaderProperties {
  children?: any;
  className?: string;
}

const Header = ({ children, className }: HeaderProperties) => {
  const toggleNavbarExpanded = useNavbarStore(
    (state) => state.toggleNavbarExpanded
  );

  return (
    <div className={className}>
      <NavbarToggle
        className="w-10 cursor-pointer"
        onClick={() => toggleNavbarExpanded()}
        icon={<MenuIcon className="h-10 w-10" />}
      />
      {children}
    </div>
  );
};

export default Header;
import { useNavigate } from "react-router-dom";
import { Icon } from "./NavbarItemIcon";
import { Title } from "./NavbarItemTitle";

interface NavbarItemProperties {
  path: string;
  className: string;
  children: any;
}

const NavbarItem = (props: NavbarItemProperties) => {
  const navigate = useNavigate();
  return (
    <div className={props.className} onClick={() => navigate(props.path)}>
      {props.children}
    </div>
  );
};

NavbarItem.Icon = Icon;
NavbarItem.Title = Title;

export default NavbarItem;

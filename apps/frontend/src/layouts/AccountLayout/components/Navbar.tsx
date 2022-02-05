import NavbarItem from "../../../components/headless/Navbar/NavbarItem";
import NavbarIcon from "../../../components/headless/Navbar/NavbarIcon";
import NavbarTitle from "../../../components/headless/Navbar/NavbarTitle";
import NavbarHeader from "../../../components/headless/Navbar/NavbarHeader";
import { Navbar as HeadlessNavbar } from "../../../components/headless/Navbar/Navbar";
import classNames from "classnames";
import { useNavbarStore } from "../../../utils/store";

export interface Route {
  path: string;
  title: string;
  icon: string;
  subRoutes?: Route[];
}

const Navbar = (items: Route[]) => {
  const navbarExpanded = useNavbarStore(state => state.navbarExpanded);

  return (
    <HeadlessNavbar className={navbarExpanded ? "w-[250px]": "hidden" }>
      <NavbarHeader
        className={classNames("flex", { "h-[50px]": navbarExpanded})}
      >
        <div className="flex h-[50px] w-full justify-center items-center">
          Navbar title
        </div>
      </NavbarHeader>
      {items.map((item) => (
        <NavbarItem
          className="flex w-full h-[40px] items-center justify-center items-center pl-8 rounded bg-gray-300"
          path={item.path}
          key={item.path}
        >
          <NavbarIcon className={classNames("w-14 items-center")} >{item.icon}</NavbarIcon>
          <NavbarTitle className={classNames("w-full items-center", {"hidden": !navbarExpanded })}>
            {item.title}
          </NavbarTitle>
        </NavbarItem>
      ))}
    </HeadlessNavbar>
  );
};

export default Navbar;

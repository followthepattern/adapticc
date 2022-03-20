import Navbar, { INavbarItem } from "./components/Navbar";
import Header from "./components/Header";

import { routes } from "../../router/routes";
import { useNavbarStore } from "../../utils/store";
import classNames from "classnames";

const AccountLayout = (props: any) => {
  const { navbarExpanded, closeNavbar } = useNavbarStore();

  const handleCloseNavbar = () => {
    if (navbarExpanded) {
      closeNavbar();
    }
  };

  const navBarItems = routes
    .filter((route) => {
      return route.showNavbar;
    })
    .map((route) => route as INavbarItem);

  return (
    <div className="flex h-screen w-screen overlay-hidden">
      <div
        className={classNames(
          "md:static bg-white xs:absolute h-full border-r",
          { hidden: !navbarExpanded }
        )}
      >
        {Navbar(navBarItems)}
      </div>
      <div className="flex flex-col w-full" onClick={handleCloseNavbar}>
        <Header className="w-full h-10" />
        <div className="grow w-full md:container mx-auto overflow-auto">{props.children}</div>
      </div>
    </div>
  );
};

export default AccountLayout;

import Navbar from "./components/Navbar";
import Header from "./components/Header";

import { mainRoutes } from "../../router/main_routes";
import { useNavbarStore } from "../../utils/store";
import classNames from "classnames";

const AccountLayout = (props: any) => {
  const {navbarExpanded, closeNavbar} = useNavbarStore();

  const handleCloseNavbar = () => {
    if (navbarExpanded) {
      closeNavbar();
    }
  }

  return (
    <div className="flex h-screen w-screen">
      <div className={classNames("md:static bg-white xs:absolute h-full border-r", {"hidden": !navbarExpanded} )}>{Navbar(mainRoutes)}</div>
      <div className="flex-col w-full" onClick={handleCloseNavbar}>
        <Header className="w-full h-10"/>
        <div className="w-full container">{props.children}</div>
      </div>
    </div>
  );
};

export default AccountLayout;

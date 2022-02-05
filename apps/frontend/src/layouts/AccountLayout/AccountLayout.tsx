import Navbar from "./components/Navbar";

import { mainRoutes } from "../../routes/main_routes";
import NavbarToggle from "../../components/headless/Navbar/NavbarToggle";
import { useNavbarStore } from "../../utils/store";

const AccountLayout = (props: any) => {
  const toggleNavbarExpanded = useNavbarStore(state => state.toggleNavbarExpanded);
  return (
    <>
      <div className="h-full flex">
        <div className="h-full border-r">
          {Navbar(mainRoutes)}
        </div>
        <div className="grow h-full">
          <div className="w-full h-[50px]">
            <NavbarToggle onClick={() => toggleNavbarExpanded()}>
              Toggle
            </NavbarToggle>
          </div>
          <div className="h-[calc(100%-50px)] w-full">
            {props.children}
          </div>
        </div>
      </div>
    </>
  );
};

export default AccountLayout;

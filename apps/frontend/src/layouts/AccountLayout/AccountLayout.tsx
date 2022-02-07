import Navbar from "./components/Navbar";
import Header from "./components/Header";

import { mainRoutes } from "../../routes/main_routes";

const AccountLayout = (props: any) => {
  return (
    <div className="flex">
      <div className="border-r">{Navbar(mainRoutes)}</div>
      <div className="h-[20px]">
        <Header className="w-full" />
        <div className="w-full">{props.children}</div>
      </div>
    </div>
  );
};

export default AccountLayout;

import Navbar from "../../components/styled/Navbar";

import { mainRoutes } from "../../routes/main_routes";

const AccountLayout = (props: any) => {
  return (
    <>
      <div className="h-full flex">
        <div className="w-[250px] h-full border-r">
          {Navbar(mainRoutes)}
        </div>
        <div className="grow h-full">
          <div className="w-full h-[50px]">header</div>
          <div className="h-[calc(100%-50px)] w-full">
            {props.children}
          </div>
        </div>
      </div>
    </>
  );
};

export default AccountLayout;

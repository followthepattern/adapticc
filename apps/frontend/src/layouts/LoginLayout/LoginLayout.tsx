import React, { FunctionComponent } from "react";

const LoginLayout = (props: any) => {
  return (
    <div className="h-screen w-screen px-4">
      <div className="flex md:h-4/5 xs:h-full w-full justify-center items-center">
        {props.children}
      </div>
    </div>
  );
};

export default LoginLayout;

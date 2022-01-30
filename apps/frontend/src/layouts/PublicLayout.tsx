import React, { FunctionComponent } from "react";

const PublicLayout = (props: any) => {
  return (
    <>
      <div>Public Layout</div>
      <div>{props.children}</div>
    </>
  );
};

export default PublicLayout;
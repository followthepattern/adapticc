import { BrowserRouter, Routes, Route } from "react-router-dom";
import WithAuthorization from "../authorization/components/WithAuthorization";

import { mainRoutes } from "./main_routes";

function Router() {
  return (
    <Routes>
      {mainRoutes.map((route) => {
        const getElement = () => (
          <route.Layout>
            <route.Page />
          </route.Layout>
        );

        const element = route.public
          ? getElement()
          : WithAuthorization(getElement);

        return <Route path={route.path} key={route.path} element={element} />;
      })}
    </Routes>
  );
}

export default Router;

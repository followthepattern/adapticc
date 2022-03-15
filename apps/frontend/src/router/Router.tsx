import { Routes, Route } from "react-router-dom";
import WithAuthorization from "../authorization/components/WithAuthorization";

import { routes } from "./routes";

function Router() {
  return (
    <Routes>
      {routes.map((route) => {
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

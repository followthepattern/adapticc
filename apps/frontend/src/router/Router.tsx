import { BrowserRouter, Routes, Route } from "react-router-dom";
import UserContext from "../contexts/UserContext";

import { mainRoutes, Route as MainRoute } from "./main_routes";

function GetPublicPage(route: MainRoute) {
  return (
    <route.Layout>
      <route.Page />
    </route.Layout>
  );
}

function GetAccountPage(route: MainRoute) {
  return (
    <route.Layout>
      <UserContext.Provider value={{
          isAuthenticated: true,
      }}>
          <route.Page/>
      </UserContext.Provider>
    </route.Layout>
  );
}

function Router() {
  return (
    <BrowserRouter>
      <Routes>
        {mainRoutes.map((route) => {
          let element;
          if (route.public) {
            element = GetPublicPage(route);
          } else {
            element = GetAccountPage(route);
          }

          return <Route path={route.path} key={route.path} element={element} />;
        })}
      </Routes>
    </BrowserRouter>
  );
}

export default Router;

import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.css";

import { mainRoutes } from "./routes/main";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {mainRoutes.map((route) => {
          return (
            <Route
              path={route.path}
              key={route.path}
              element={<route.Layout>
                <route.Page/>
              </route.Layout>}
            />
          );
        })}
      </Routes>
    </BrowserRouter>
  );
}

export default App;

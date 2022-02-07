import { BrowserRouter, Routes, Route } from "react-router-dom";
import {
  ApolloClient,
  ApolloProvider,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import "./App.css";

import { mainRoutes } from "./routes/main_routes";

function App() {
  const httpLink = createHttpLink({
    uri: "http://localhost:3011/graphql",
  });

  const authLink = setContext((_, { headers }) => {
    // get the authentication token from local storage if it exists
    // const token = localStorage.getItem("token");
    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        Authorization: `Bearer test3`,
      },
    };
  });

  const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <Routes>
          {mainRoutes.map((route) => {
            return (
              <Route
                path={route.path}
                key={route.path}
                element={
                  <route.Layout>
                    <route.Page />
                  </route.Layout>
                }
              />
            );
          })}
        </Routes>
      </BrowserRouter>
    </ApolloProvider>
  );
}

export default App;

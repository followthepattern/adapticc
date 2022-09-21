import {
  ApolloClient,
  ApolloProvider,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { BrowserRouter } from "react-router-dom";

import Router from "./router/Router";
import { GetTokenFromStorage } from "./utils/store";

function App() {
  const jwt = GetTokenFromStorage();

  const httpLink = createHttpLink({
    uri: "http://localhost:3011/graphql",
  });

  const authLink = setContext((_, { headers }) => {
    return {
      headers: {
        ...headers,
        Authorization: jwt ? `Bearer ${jwt}` : "",
      },
    };
  });

  console.debug("App render");

  const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <Router />
      </BrowserRouter>
    </ApolloProvider>
  );
}

export default App;

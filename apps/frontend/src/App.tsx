import {
  ApolloClient,
  ApolloProvider,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

import Router from "./router/Router";
import { GetTokenFromStorage } from "./utils/store";

function App() {
  const jwt = GetTokenFromStorage();
  console.info("jwt:", jwt);

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

  console.debug("APP RENDER!");

  const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  return (
    <ApolloProvider client={client}>
      <Router />
    </ApolloProvider>
  );
}

export default App;

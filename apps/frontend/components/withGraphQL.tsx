import { ApolloClient, ApolloProvider, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { API_GRAPHQL } from "@/lib/config";

interface WithGraphQLProperties {
    children?: any;
    token?: string;
}

const WithGraphQL = ({children, token}: WithGraphQLProperties) => {
    const httpLink = createHttpLink({
        uri: API_GRAPHQL,
    });

    const authorizationHeaderText = token ? `Bearer ${token}` : "";

    const authLink = setContext((_, { headers }) => {
        return {
            headers: {
                ...headers,
                Authorization: authorizationHeaderText,
            },
        };
    });

    const client = new ApolloClient({
        link: authLink.concat(httpLink),
        cache: new InMemoryCache(),
    })

    return (
        <ApolloProvider client={client}>
            {children}
        </ApolloProvider>
    )
}

export default WithGraphQL;
import { ApolloClient, ApolloProvider, InMemoryCache, createHttpLink } from "@apollo/client";
import { RetryLink } from "@apollo/client/link/retry";
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
    }).concat(httpLink);

    const retryLink = new RetryLink({
        attempts: (count, operation, error) => {
            return !!error && operation.operationName != 'specialCase';
        },
        delay: (count, operation, error) => {
            return count * 1000 * Math.random();
        },
    }).concat(authLink) ;

    const client = new ApolloClient({
        link: retryLink,
        cache: new InMemoryCache(),
    })

    return (
        <ApolloProvider client={client}>
            {children}
        </ApolloProvider>
    )
}

export default WithGraphQL;
package graphql_api

const Schema = `
scalar Time

type User {
	id: String
	email: String
	firstName: String
	lastName: String
	active: Boolean
	lastLoginAt: Time
	createdAt: Time
	creationUserID: String
	updatedAt: Time
	updateUserID: String
}

type LoginResponse {
	jwt: String
	expires_at: Time
}

type RegisterResponse {
	email: String
	first_name: String
	last_name: String
}

schema {
	query: Query
	mutation: Mutation
}

type Query {
	users: UserQuery!
}

type Mutation {
	authentication: AuthMutation!
}

type UserQuery {
	single(id: String!): User
	profile: User
}

type AuthMutation {
	login(email: String!, password: String!): LoginResponse
	register(email: String!, firstName: String!, lastName: String!, password: String!): RegisterResponse
}
`

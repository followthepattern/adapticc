package graphql_api

const Schema = `
scalar Time
scalar Int64
scalar Uint

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

input UserListFilter {
	pageSize: Uint
	page: Uint
	name: String
	email: String
}

type UserListResponse {
	count: Int64!
	pageSize: Uint
	page: Uint
	data: [User]!
}

type LoginResponse {
	jwt: String
	expires_at: Time
}

type RegisterResponse {
	email: String
	firstName: String
	lastName: String
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
	list(filter: UserListFilter!): UserListResponse
	profile: User
}

type AuthMutation {
	login(email: String!, password: String!): LoginResponse
	register(email: String!, firstName: String!, lastName: String!, password: String!): RegisterResponse
}
`

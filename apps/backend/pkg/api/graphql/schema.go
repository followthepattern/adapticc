package graphql

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
	registeredAt: Time
}

input UserListFilter {
	pageSize: Uint
	page: Uint
	search: String
}

type UserListResponse {
	count: Int64!
	pageSize: Uint
	page: Uint
	data: [User]!
}

type Product {
	id: String
	title: String
	description: String
}

input ProductInput {
	id: String
	title: String
	description: String
}

input ProductListFilter {
	pageSize: Uint
	page: Uint
	search: String
	id: String
}

type ProductListResponse {
	count: Int64!
	pageSize: Uint
	page: Uint
	data: [Product!]!
}

type LoginResponse {
	jwt: String!
	expires_at: Time!
}

type RegisterResponse {
	email: String
	first_name: String
	last_name: String
}

type ResponseStatus {
	code: Uint!
}

schema {
	query: Query
	mutation: Mutation
}

type Query {
	users: UserQuery!
	products: ProductQuery!
}

type Mutation {
	authentication: AuthMutation!
	users: UserMutation!
	products: ProductMutation!
}

type UserQuery {
	single(id: String!): User
	list(filter: UserListFilter!): UserListResponse
	profile: User
}

type ProductQuery {
	single(id: String!): Product
	list(filter: ProductListFilter!): ProductListResponse
}

type UserMutation {
	update(id: String!, firstName: String, lastName: String): ResponseStatus
	delete(id: String!): ResponseStatus

}

type ProductMutation {
	create(model: ProductInput!): ResponseStatus
	update(model: ProductInput!): ResponseStatus
	delete(id: String!): ResponseStatus
}

type AuthMutation {
	login(email: String!, password: String!): LoginResponse
	register(email: String!, firstName: String!, lastName: String!, password: String!): RegisterResponse
}
`

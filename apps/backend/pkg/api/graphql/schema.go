package graphql

const Schema = `
scalar Time
scalar Int64
scalar Uint

type ResponseStatus {
	code: Uint!
}

input Pagination {
	pageSize: Uint
	page: Uint
}

input OrderBy {
	name: String!
	desc: Boolean
}

input ListFilter {
	search: String
}

type User {
	id: String
	email: String
	firstName: String
	lastName: String
	active: Boolean
	registeredAt: Time
}

input UserInput {
	email: String
	firstName: String
	lastName: String
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
	title: String
	description: String
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
	list(pagination: Pagination, filter: ListFilter, orderBy: [OrderBy!]): UserListResponse
	profile: User
}

type ProductQuery {
	single(id: String!): Product
	list(pagination: Pagination!, filter: ListFilter, orderBy: [OrderBy!]): ProductListResponse
}

type UserMutation {
	create(model: UserInput!): ResponseStatus
	update(id: String!, model: UserInput!): ResponseStatus
	delete(id: String!): ResponseStatus

}

type ProductMutation {
	create(model: ProductInput!): ResponseStatus
	update(id: String!, model: ProductInput!): ResponseStatus
	delete(id: String!): ResponseStatus
}

type AuthMutation {
	login(email: String!, password: String!): LoginResponse
	register(email: String!, firstName: String!, lastName: String!, password: String!): RegisterResponse
}
`

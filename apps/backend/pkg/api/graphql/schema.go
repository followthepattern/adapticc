package graphql

const Schema = `
scalar Time
scalar NullInt64
scalar NullUint
scalar NullInt

type ResponseStatus {
	code: Int!
}

input Pagination {
	pageSize: NullUint!
	page: NullUint!
}

input OrderBy {
	name: String!
	desc: Boolean
}

input ListFilter {
	search: String!
}

type User {
	id: String!
	email: String!
	firstName: String!
	lastName: String!
	active: Boolean!
	creationUserID: String!
	updateUserID: String!
	createdAt: Time!
	updatedAt: Time!
}

input UserInput {
	email: String!
	firstName: String!
	lastName: String!
}

type UserListResponse {
	count: NullInt64!
	pageSize: NullUint!
	page: NullUint!
	data: [User!]!
}

type Product {
	id: String!
	title: String!
	description: String!
	creationUserID: String!
	updateUserID: String!
	createdAt: Time!
	updatedAt: Time!
}

input ProductInput {
	title: String!
	description: String!
}

type ProductListResponse {
	count: NullInt64!
	pageSize: NullUint!
	page: NullUint!
	data: [Product!]!
}

type Role {
	id: String!
	code: String!
	name: String!
	creationUserID: String!
	updateUserID: String!
	createdAt: Time!
	updatedAt: Time!
}

type RoleListResponse {
	count: NullInt64!
	pageSize: NullUint!
	page: NullUint!
	data: [Role!]!
}

type LoginResponse {
	jwt: String!
	expires_at: Time!
}

type RegisterResponse {
	email: String!
	first_name: String!
	last_name: String!
}

schema {
	query: Query
	mutation: Mutation
}

type Query {
	users: UserQuery!
	products: ProductQuery!
	roles: RoleQuery!
}

type Mutation {
	authentication: AuthMutation!
	users: UserMutation!
	products: ProductMutation!
	roles: RoleMutation!
}

type UserQuery {
	single(id: String!): User
	list(pagination: Pagination, filter: ListFilter, orderBy: [OrderBy!]): UserListResponse
	profile: User
}

type ProductQuery {
	single(id: String!): Product
	list(pagination: Pagination, filter: ListFilter, orderBy: [OrderBy!]): ProductListResponse
}

type RoleQuery {
	single(id: String!): Role
	list(pagination: Pagination, filter: ListFilter, orderBy: [OrderBy!]): RoleListResponse
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

type RoleMutation {
	addRoleToUser(userID: String!, roleID: String!): ResponseStatus
	deleteRoleFromUser(userID: String!, roleID: String!): ResponseStatus
}
`
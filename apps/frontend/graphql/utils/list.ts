export interface List<T> {
	count?: number
	page?: number
	pageSize?: number
	data?: T[]
}
export interface ListFilter {
	page?: number
	pageSize?: number
	search?: string
}

export interface List<T> {
	count?: number
	page?: number
	pageSize?: number
	search?: string
	data?: T[]
}
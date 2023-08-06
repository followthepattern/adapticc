import { User } from "@/models/user";
import { Product } from "@/models/product";
import { ListResponse } from "./utils/list";

export interface SingleQueryResult<Data = any> {
    loading: boolean;
    data?: Data;
    error?: any;
    itemNotFound?: boolean;
}

export interface QueryResourceResponse<T> {
	profile?: T
	single?: T
	list?: ListResponse<T>
}

export interface QueryResponse {
	users?: QueryResourceResponse<User>
	products?: QueryResourceResponse<Product>
}
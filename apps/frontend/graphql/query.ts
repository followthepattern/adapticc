import { Products } from "./products/query";
import { Users } from "./users/query";

export interface QueryResponse {
	users?: Users
	products?: Products
}
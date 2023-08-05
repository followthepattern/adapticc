import { ProductMutation } from "./products/mutation";
import { Users } from "./users/mutation";

export interface MutationResponse {
    users: Users
    products: ProductMutation
}
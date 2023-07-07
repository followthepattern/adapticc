import { User } from "@/graphql/users/query";
import { createContext } from "react";

const UserContext = createContext<User>({});

export default UserContext;

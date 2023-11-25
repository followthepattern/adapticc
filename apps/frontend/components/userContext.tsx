import { User } from "@/models/user";
import { createContext } from "react";

const UserContext = createContext<User>({id:""});

export default UserContext;
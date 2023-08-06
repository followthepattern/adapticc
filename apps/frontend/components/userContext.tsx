import { User } from "@/models/user";
import { createContext } from "react";

const UserContext = createContext<User>({});

export default UserContext;

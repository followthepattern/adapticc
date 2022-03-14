import { useContext } from "react";
import UserContext from "../../../contexts/UserContext";

const Dashboard = (props: any) => {
  const userProfile = useContext(UserContext);

  return <p className="w-full break-words">{JSON.stringify(userProfile)}</p>;
};

export default Dashboard;

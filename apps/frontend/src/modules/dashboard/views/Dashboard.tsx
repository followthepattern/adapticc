import { gql, useLazyQuery } from "@apollo/client";

import { getUser } from "../../../graphql/user/query";

const Dashboard = (props: any) => {
  const [loadGreeting, { called, loading, data, error }] = useLazyQuery(
    gql(getUser),
    { variables: { language: "english" } }
  );

  if (error) return <p>{JSON.stringify(error)}</p>;

  if (called && loading) return <p>loading</p>;
  if (!called) {
    loadGreeting();
  }

  return <p className="w-full break-words">{JSON.stringify(data?.users?.single)}</p>;
};

export default Dashboard;

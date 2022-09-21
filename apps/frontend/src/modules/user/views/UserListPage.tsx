import { useLazyQuery, gql } from "@apollo/client";
import { useEffect } from "react";
import List, {ListProperties} from "../../../components/styled/List";
import { getUsers } from "../../../graphql/user/query";

const Users = (props: any) => {
  const [executeGetUsers, {data, called, loading, error}] = useLazyQuery(gql(getUsers))

  useEffect(() => {
    if (!called) {
      executeGetUsers({
        variables: {
          page: 2,
          pageSize: 3,
          search: null,
        }
      });
    }
  }, [])

  if (!called || loading) {
    console.info("loading")
    return <div>Loading...</div>
  }

  if (error) {
    return <div>{JSON.stringify(error)}</div>
  }

  console.info(data);

  let listProperties : ListProperties = {
    columns: ["id", "email", "First Name", "Last Name", ""],
    rows: [],
    paginationData: {
      current: data.users.list.page,
      first: 1,
      last: Math.ceil(data.users.list.count / data.users.list.pageSize),
      limit: data.users.list.pageSize,
    },
    resourceName: "users",
  }

  data.users.list.data.map((record: any)=> {
    listProperties.rows.push({
      cells: [record.id, record.email, record.firstName, record.lastName],
      action: "Edit"
    })
  })

  return (
    <>
      <div className="inline-block text-3xl px-1 py-4">Users</div>
      {List(listProperties)}
    </>
  );
};

export default Users;

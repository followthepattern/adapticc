import UserContext from "@/components/userContext";
import { useContext } from "react";
import SingleLayout from "../components/singleView/layout";
import DataListView from "../components/singleView/dataListView/dataListView";

export default function Profile() {
  const userProfile = useContext(UserContext);

  return (
    <SingleLayout>
      <SingleLayout.Title>Profile</SingleLayout.Title>
      <SingleLayout.Subtitle>Personal settings</SingleLayout.Subtitle>
      <DataListView className="my-6">
        <DataListView.Row>
          <DataListView.Label>
            Full Name
          </DataListView.Label>
          <DataListView.Field>
            {userProfile?.firstName} {userProfile?.lastName}
          </DataListView.Field>
        </DataListView.Row>
        <DataListView.Row>
          <DataListView.Label>
            Email
          </DataListView.Label>
          <DataListView.Field>
            {userProfile?.email}
          </DataListView.Field>
        </DataListView.Row>
      </DataListView>
    </SingleLayout>
  )
}
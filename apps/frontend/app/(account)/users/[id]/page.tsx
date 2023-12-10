import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../page";
import useSingleUser from "../hooks/singleUser";
import SingleLayout from "../../components/singleView/layout";
import DataListView from "../../components/singleView/dataListView/dataListView";
import SecondaryButton from "../../components/buttons/secondaryButton";
import PrimaryButton from "../../components/buttons/primaryButton";

export default function User() {
    const { id } = useParams()
    const navigate = useNavigate();

    const { loading, data, error, itemNotFound } = useSingleUser(id ?? "");

    if (loading) {
        return (
            <div>Loading...</div>
        )
    }

    if (error) {
        return (
            <div>{JSON.stringify(error)}</div>
        )
    }

    if (itemNotFound) {
        // notFound();
    }

    const onCancel = () => {
        navigate(RESOURCE_URL);
    }

    const onEdit = () => {
        navigate(`${RESOURCE_URL}/${id}/edit`);
    }

    return (
        <SingleLayout>
            <SingleLayout.Title>User</SingleLayout.Title>
            <DataListView className="mt-6">
                <DataListView.Row>
                    <DataListView.Label>
                        ID
                    </DataListView.Label>
                    <DataListView.Field>
                        {data?.id}
                    </DataListView.Field>
                </DataListView.Row>
                <DataListView.Row>
                    <DataListView.Label>
                        Email
                    </DataListView.Label>
                    <DataListView.Field>
                        {data?.email}
                    </DataListView.Field>
                </DataListView.Row>
                <DataListView.Row>
                    <DataListView.Label>
                        First Name
                    </DataListView.Label>
                    <DataListView.Field>
                        {data?.firstName}
                    </DataListView.Field>
                </DataListView.Row>
                <DataListView.Row>
                    <DataListView.Label>
                        Last Name
                    </DataListView.Label>
                    <DataListView.Field>
                        {data?.lastName}
                    </DataListView.Field>
                </DataListView.Row>
            </DataListView>
            <SingleLayout.Footer className="justify-end">
                <SecondaryButton onClick={onCancel}>
                    Cancel
                </SecondaryButton>
                <PrimaryButton onClick={onEdit}>
                    Edit
                </PrimaryButton>
            </SingleLayout.Footer>
        </SingleLayout>
    )
}
import { useNavigate, useParams } from "react-router-dom";
import { RESOURCE_URL } from "../page";
import useSingleUser from "../hooks/singleUser";
import SingleLayout from "../../components/singlePage/layout";
import DataList from "../../components/dataList/dataList";
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
            <DataList className="mt-6">
                <DataList.Row name={"ID"}>
                    {data?.id}
                </DataList.Row>
                <DataList.Row name={"Email"}>
                    {data?.email}
                </DataList.Row>
                <DataList.Row name={"First Name"}>
                    {data?.firstName}
                </DataList.Row>
                <DataList.Row name={"Last Name"}>
                    {data?.lastName}
                </DataList.Row>
            </DataList>
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
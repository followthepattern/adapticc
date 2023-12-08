import { useNavigate, useParams } from "react-router-dom";
import useSingleProduct from "../hooks/singleProduct";
import { RESOURCE_URL } from "../page";
import SingleLayout from "../../components/singlePage/layout";
import DataListView from "../../components/DataListView/DataListView";
import SecondaryButton from "../../components/buttons/secondaryButton";
import PrimaryButton from "../../components/buttons/primaryButton";

export default function Product() {
    const { id } = useParams()
    const navigate = useNavigate();

    const { loading, data, error, itemNotFound } = useSingleProduct(id ?? "");

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
            <SingleLayout.Title>Product</SingleLayout.Title>
            <DataListView className="mt-6">
                <DataListView.Row name={"ID"}>
                    {data?.id}
                </DataListView.Row>
                <DataListView.Row name={"Title"}>
                    {data?.title}
                </DataListView.Row>
                <DataListView.Row name={"Description"}>
                    {data?.description}
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
import { useNavigate, useParams } from "react-router-dom";
import useSingleProduct from "../hooks/singleProduct";
import { RESOURCE_URL } from "../page";
import SingleLayout from "../../components/singlePage/layout";
import DataList from "../../components/dataList/dataList";
import CancelButton from "../../components/buttons/cancelButton";
import EditButton from "../../components/buttons/editButton";

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
            <DataList className="mt-6">
                <DataList.Row name={"ID"}>
                    {data?.id}
                </DataList.Row>
                <DataList.Row name={"Title"}>
                    {data?.title}
                </DataList.Row>
                <DataList.Row name={"Description"}>
                    {data?.description}
                </DataList.Row>
            </DataList>
            <SingleLayout.Footer>
                <CancelButton onClick={onCancel} />
                <EditButton onClick={onEdit} />
            </SingleLayout.Footer>
        </SingleLayout>
    )
}
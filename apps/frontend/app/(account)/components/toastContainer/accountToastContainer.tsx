import { ToastContainer, TypeOptions } from "react-toastify";

function toastClassNameSelector(type?: TypeOptions): string {
    return "relative flex p-4 bg-white m-1 min-h-10 text-gray-900 rounded-lg shadow-lg justify-between overflow-hidden cursor-pointer"
}

export default function AccountToastContainer() {
    return (
        <ToastContainer
            toastClassName={(context) => toastClassNameSelector(context?.type)}
            bodyClassName={() => "flex text-sm font-medium p-3"}
            position="top-right"
            autoClose={3000}
            hideProgressBar={true}
            newestOnTop={false}
            closeOnClick
            rtl={false}
            pauseOnFocusLoss={false}
            draggable
            pauseOnHover
        />
    )
}
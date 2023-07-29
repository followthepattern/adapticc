import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import { Routes } from './routes';

function App() {
    const router = createBrowserRouter(Routes);

    return (
        <RouterProvider router={router} />
    );
};

export default App;
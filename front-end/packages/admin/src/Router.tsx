import {
    // BrowserRouter,
    Routes,
    Route,
} from "react-router-dom";
import { Collections } from "./containers/Collections";
import { Home } from "./containers/Home";

export const AppRouter = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/collections" element={<Collections />} />
        </Routes>
    )
}
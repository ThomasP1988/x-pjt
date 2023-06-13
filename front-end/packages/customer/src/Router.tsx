import {
    // BrowserRouter,
    Routes,
    Route,
} from "react-router-dom";
import { Deposit } from "./containers/Deposit";
import { DepositNFT } from "./containers/DepositNFT";
import { DepositNFTCheckOut } from "./containers/DepositNFTCheckOut";
import { Discover } from "./containers/Discover";
import { Home } from "./containers/Home";
import { Market } from "./containers/Market";
import { Withdrawal } from "./containers/Withdrawal";


export const AppRouter = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/deposit" element={<Deposit />} />
            <Route path="/deposit/nft" element={<DepositNFT />} />
            <Route path="/deposit/nft/checkout" element={<DepositNFTCheckOut />} />
            <Route path="/withdrawal" element={<Withdrawal />} />
            <Route path="/discover" element={<Discover />} />
            <Route path="/m/:symbol" element={<Market />} />
            <Route path="/m/:symbol/:tab" element={<Market />} />
        </Routes>
    )
}
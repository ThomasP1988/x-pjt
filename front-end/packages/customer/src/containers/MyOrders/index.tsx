import Container from '@mui/material/Container';
import { OrderList } from 'shared/components/OrderList';

type Props = {
    symbol?: string
}

export const MyOrders = (args: Props) => {
    return (
        <Container sx={{ marginTop: -7 }}>
           <OrderList {...args} isOpen={true} />
        </Container>
    )
}
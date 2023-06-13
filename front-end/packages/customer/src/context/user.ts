import { useContext, createContext } from "react";

type Context = {
  user: any;
  openSignInDialog: () => void
}

export const UserContext = createContext<Context>({
  user: null,
  openSignInDialog: () => {},
});

export function useUserContext(): Context {
  return useContext(UserContext);
}
import { Dispatch, SetStateAction, createContext, useContext } from "react";

export type User = {
  userName: string;
};

export const UserContext = createContext<
  [User, Dispatch<SetStateAction<User>>] | undefined
>(undefined);

// export const UserContext = createContext<any>(undefined);

export const useUserContext = () => {
  const user = useContext(UserContext);

  if (user === undefined) {
    throw new Error("useUserContext is undefined");
  }

  return user;
};

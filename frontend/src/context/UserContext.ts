import { Dispatch, SetStateAction, createContext, useContext } from "react";

export type User = {
  userName: string;
};

type UserContextState = [user: User, setUser: Dispatch<SetStateAction<User>>];

export const UserContext = createContext<UserContextState | undefined>(
  undefined
);

// export const UserContext = createContext<any>(undefined);

export const useUserContext = () => {
  const user = useContext(UserContext);

  if (user === undefined) {
    throw new Error("useUserContext is undefined");
  }

  return user;
};

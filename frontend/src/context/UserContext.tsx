import {
  Dispatch,
  SetStateAction,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import { User } from "../types";
import { fetchUser } from "../api/getUser";

type IUserContext = [user: User, setUser: Dispatch<SetStateAction<User>>];

export const UserContext = createContext<IUserContext | undefined>(undefined);

export const useUserContext = (): IUserContext => {
  const user = useContext(UserContext);

  if (user === undefined) {
    throw new Error("useUserContext is undefined");
  }

  return user;
};

export const UserContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const defaultUser: User = {
    userName: "",
  };

  const [user, setUser] = useState<User>(defaultUser);

  useEffect(() => {
    fetchUser().then((user) => {
      if (user) {
        setUser(user);
      }
    });
  }, []);

  return (
    <UserContext.Provider value={[user, setUser]}>
      {children}
    </UserContext.Provider>
  );
};
